// File: internal/domain/table/tableentity.go
package tableentity

type ColumnType string

const (
	TypeVARCHAR   ColumnType = "varchar"
	TypeTEXT      ColumnType = "text"
	TypeINT       ColumnType = "integer"
	TypeBIGINT    ColumnType = "bigint"
	TypeFLOAT     ColumnType = "float"
	TypeDOUBLE    ColumnType = "double precision"
	TypeBOOLEAN   ColumnType = "boolean"
	TypeDATE      ColumnType = "date"
	TypeTIMESTAMP ColumnType = "timestamp"
	TypeJSON      ColumnType = "jsonb"
)

type Column struct {
	Name          string     `json:"name"`
	Type          ColumnType `json:"type"`
	Length        *int       `json:"length,omitempty"`
	NotNull       bool       `json:"not_null"`
	PrimaryKey    bool       `json:"primary_key"`
	AutoIncrement bool       `json:"auto_increment"`
	Unique        bool       `json:"unique"`
	Default       *string    `json:"default,omitempty"`
}

type Table struct {
	Name        string   `json:"name"`
	Columns     []Column `json:"columns"`
	Description string   `json:"description,omitempty"`
}
