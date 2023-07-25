package schema

import (
	"go/ast"
	"reflect"

	"github.com/go-labx/orm/dialect"
)

// Field represents a column of database
type Field struct {
	Name string // Name is the name of the field.
	Type string // Type is the type of the field.
	Tag  string // Tag is the tag associated with the field.
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}       // Model is the model of the schema.
	Name       string            // Name is the name of the schema.
	Fields     []*Field          // Fields is a slice of pointers to the fields in the schema.
	FieldNames []string          // FieldNames is a slice of the names of the fields in the schema.
	fieldMap   map[string]*Field // fieldMap is a map with field names as keys and pointers to the fields as values.
}

// GetField returns a pointer to the field with the given name in the schema.
func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

// Parse is a function that takes a destination interface and a dialect, and returns a pointer to a Schema.
// It first gets the type of the model from the destination interface, and initializes a new Schema with the model, its name, and an empty fieldMap.
// Then, it iterates over the fields of the model. If a field is not anonymous and is exported, it creates a new Field with the name of the model field and its type, as determined by the dialect.
// If the model field has a tag "orm", it also sets the Tag of the Field to the value of this tag.
// Finally, it adds the new Field to the Fields slice of the Schema, its name to the FieldNames slice, and a mapping from its name to the Field itself to the fieldMap.
// After all fields have been processed, it returns the pointer to the Schema.
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
