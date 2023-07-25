package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type MysqlDialect struct {
}

func init() {
	RegisterDialect("mysql", &MysqlDialect{})
}

func (mysql *MysqlDialect) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "BOOL"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "INT"
	case reflect.Int64, reflect.Uint64:
		return "BIGINT"
	case reflect.Float32:
		return "FLOAT"
	case reflect.Float64:
		return "DOUBLE"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (mysql *MysqlDialect) IsTableExist(tableName string) (string, []interface{}) {
	return "SELECT table_name FROM information_schema.tables WHERE table_name = ?", []interface{}{tableName}
}
