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

func main() {
	domainDir := "internal/domain"
	models, err := parseModelsFromDomain(domainDir)
	if err != nil {
		log.Fatalf("Error parsing models: %v", err)
	}

	for _, model := range models {
		generateMigration(model)
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
			if x.Name.Name == modelName {
				if structType, ok := x.Type.(*ast.StructType); ok {
					model = &Model{Name: modelName}
					for _, field := range structType.Fields.List {
						if len(field.Names) > 0 {
							fieldName := field.Names[0].Name
							fieldType := fmt.Sprintf("%s", field.Type)
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

func generateMigration(model Model) {
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_create_%s.sql", timestamp, strings.ToLower(model.Name))

	content := fmt.Sprintf("-- Create %s table\n", model.Name)
	content += fmt.Sprintf("CREATE TABLE %s (\n", strings.ToLower(model.Name))

	for i, field := range model.Fields {
		sqlType := getSQLType(field.Type)
		constraints := getConstraints(field)

		content += fmt.Sprintf("    %s %s%s", strings.ToLower(field.Name), sqlType, constraints)
		if i < len(model.Fields)-1 {
			content += ","
		}
		content += "\n"
	}

	content += ");\n"

	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Error writing migration file: %v", err)
	}

	fmt.Printf("Generated migration file: %s\n", fileName)
}

func getSQLType(goType string) string {
	switch goType {
	case "string":
		return "VARCHAR(255)"
	case "int", "int32", "int64", "uint", "uint32", "uint64":
		return "INTEGER"
	case "float32", "float64":
		return "DECIMAL(10, 2)"
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

	if field.Name == "ID" {
		constraints = append(constraints, "PRIMARY KEY")
		constraints = append(constraints, "AUTOINCREMENT")
	}

	if _, ok := field.Tags["json"]; ok && field.Tags["json"] != "-" {
		constraints = append(constraints, "NOT NULL")
	}

	if field.Name == "Email" {
		constraints = append(constraints, "UNIQUE")
	}

	if len(constraints) > 0 {
		return " " + strings.Join(constraints, " ")
	}
	return ""
}
