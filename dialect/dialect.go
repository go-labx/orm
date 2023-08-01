// Package dialect provides an interface for different SQL dialects
package dialect

import "reflect"

// dialectsMap stores all registered SQL dialects
var dialectsMap = map[string]Dialect{}

// Dialect is an interface that represents a SQL dialect
type Dialect interface {
	VersionSQL() string

	// DataTypeOf returns the SQL data type of the given Go data type
	DataTypeOf(typ reflect.Value) string
	// IsTableExistSQL returns the SQL query that checks if a table exists
	IsTableExistSQL(tableName string) string

	DropTableSQL(tableName string) string
}

// RegisterDialect registers a new SQL dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect retrieves a registered SQL dialect
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
