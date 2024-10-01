package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Model struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Tags map[string]string
}

type Index struct {
	Name    string
	Columns []string
	Unique  bool
}

func main() {
	rootDir, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	domainDir := filepath.Join(rootDir, "internal", "domain")
	migrationsDir := filepath.Join(rootDir, "migrations")
	err = os.MkdirAll(migrationsDir, 0755)
	if err != nil {
		log.Fatalf("Error creating migrations directory: %v", err)
	}

	models, err := parseModelsFromDomain(domainDir)
	if err != nil {
		log.Fatalf("Error parsing models: %v", err)
	}

	for _, model := range models {
		err := generateMigration(migrationsDir, model)
		if err != nil {
			log.Fatalf("Error generating migration: %v", err)
		}
	}
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root")
		}
		dir = parent
	}
}

func parseModelsFromDomain(domainDir string) ([]Model, error) {
	var models []Model

	err := filepath.Walk(domainDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			modelName := strings.TrimSuffix(info.Name(), ".go")
			model, err := parseModelFromFile(path, modelName)
			if err != nil {
				return err
			}
			if model != nil {
				models = append(models, *model)
			}
		}
		return nil
	})

	return models, err
}

func parseModelFromFile(filePath, modelName string) (*Model, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var model *Model

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if strings.ToLower(x.Name.Name) == strings.ToLower(modelName) {
				if structType, ok := x.Type.(*ast.StructType); ok {
					model = &Model{Name: modelName}
					for _, field := range structType.Fields.List {
						if len(field.Names) > 0 {
							fieldName := field.Names[0].Name
							fieldType := getTypeString(field.Type)
							tags := parseTags(field.Tag)
							model.Fields = append(model.Fields, Field{Name: fieldName, Type: fieldType, Tags: tags})
						}
					}
				}
			}
		}
		return true
	})

	return model, nil
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getTypeString(t.X), t.Sel.Name)
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getTypeString(t.Key), getTypeString(t.Value))
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return fmt.Sprintf("%T", expr)
	}
}

func parseTags(tag *ast.BasicLit) map[string]string {
	tags := make(map[string]string)
	if tag != nil {
		tagStr := strings.Trim(tag.Value, "`")
		for _, t := range strings.Split(tagStr, " ") {
			parts := strings.SplitN(t, ":", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := strings.Trim(parts[1], "\"")
				tags[key] = value
			}
		}
	}
	return tags
}

func generateMigration(migrationsDir string, model Model) error {
	timestamp := time.Now().Format("20060102150405")
	baseName := fmt.Sprintf("%s_create_%s", timestamp, strings.ToLower(model.Name))
	upFileName := filepath.Join(migrationsDir, baseName+".up.sql")
	downFileName := filepath.Join(migrationsDir, baseName+".down.sql")

	upContent := generateUpSQL(model)
	downContent := generateDownSQL(model)

	err := os.WriteFile(upFileName, []byte(upContent), 0644)
	if err != nil {
		return fmt.Errorf("Error writing up migration file: %v", err)
	}

	err = os.WriteFile(downFileName, []byte(downContent), 0644)
	if err != nil {
		return fmt.Errorf("Error writing down migration file: %v", err)
	}

	fmt.Printf("Generated migration files: \n%s\n%s\n", upFileName, downFileName)
	return nil
}

func generateUpSQL(model Model) string {
	content := fmt.Sprintf("-- Create %s table\n", model.Name)
	content += fmt.Sprintf("CREATE TABLE %s (\n", strings.ToLower(model.Name))

	var indexes []Index

	for i, field := range model.Fields {
		sqlType := getSQLType(field.Type)
		constraints := getConstraints(field)

		content += fmt.Sprintf("    %s %s%s", strings.ToLower(field.Name), sqlType, constraints)
		if i < len(model.Fields)-1 {
			content += ","
		}
		content += "\n"

		if dbTag, ok := field.Tags["db"]; ok {
			tags := strings.Split(dbTag, ",")
			for _, tag := range tags {
				if tag == "index" || tag == "unique_index" {
					indexes = append(indexes, Index{
						Name:    fmt.Sprintf("idx_%s_%s", strings.ToLower(model.Name), strings.ToLower(field.Name)),
						Columns: []string{strings.ToLower(field.Name)},
						Unique:  tag == "unique_index",
					})
				}
			}
		}
	}

	content += ");\n"

	// Create indexes
	for _, index := range indexes {
		if index.Unique {
			content += fmt.Sprintf("\nCREATE UNIQUE INDEX %s ON %s (%s);\n",
				index.Name, strings.ToLower(model.Name), strings.Join(index.Columns, ", "))
		} else {
			content += fmt.Sprintf("\nCREATE INDEX %s ON %s (%s);\n",
				index.Name, strings.ToLower(model.Name), strings.Join(index.Columns, ", "))
		}
	}

	return content
}

func generateDownSQL(model Model) string {
	content := fmt.Sprintf("-- Drop %s table\n", model.Name)
	content += fmt.Sprintf("DROP TABLE IF EXISTS %s;\n", strings.ToLower(model.Name))

	// Drop indexes
	for _, field := range model.Fields {
		if dbTag, ok := field.Tags["db"]; ok {
			tags := strings.Split(dbTag, ",")
			for _, tag := range tags {
				if tag == "index" || tag == "unique_index" {
					indexName := fmt.Sprintf("idx_%s_%s", strings.ToLower(model.Name), strings.ToLower(field.Name))
					content += fmt.Sprintf("\nDROP INDEX IF EXISTS %s;\n", indexName)
				}
			}
		}
	}

	return content
}

func getSQLType(goType string) string {
	if strings.HasPrefix(goType, "*") {
		goType = strings.TrimPrefix(goType, "*")
	}

	switch goType {
	case "string":
		return "VARCHAR(255)"
	case "int", "int8", "int16", "int32", "int64":
		return "INTEGER"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "INTEGER"
	case "float32", "float64":
		return "REAL"
	case "bool":
		return "BOOLEAN"
	case "time.Time":
		return "TIMESTAMP"
	default:
		return "TEXT"
	}
}

func getConstraints(field Field) string {
	var constraints []string

	if dbTag, ok := field.Tags["db"]; ok {
		tags := strings.Split(dbTag, ",")
		for _, tag := range tags {
			switch tag {
			case "primary_key":
				constraints = append(constraints, "PRIMARY KEY")
			case "auto_increment":
				constraints = append(constraints, "AUTOINCREMENT")
			case "not_null":
				constraints = append(constraints, "NOT NULL")
			case "unique":
				constraints = append(constraints, "UNIQUE")
			}
		}
	} else {
		// Add NOT NULL if it's not a pointer type
		if !strings.HasPrefix(field.Type, "*") {
			constraints = append(constraints, "NOT NULL")
		}
	}

	if len(constraints) > 0 {
		return " " + strings.Join(constraints, " ")
	}
	return ""
}
