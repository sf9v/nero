package nero

import (
	"fmt"
	"reflect"
	"text/template"

	"github.com/sf9v/mira"
)

// Templater is an interface that wraps the Filename and Template method
type Templater interface {
	// Filename is the filename of the generated file
	Filename() string
	// Template is template for generating the repository implementation
	Template() string
}

// ParseTemplater parses the repository templater
func ParseTemplater(tmpl Templater) (*template.Template, error) {
	return template.New(tmpl.Filename() + ".tmpl").
		Funcs(NewFuncMap()).Parse(tmpl.Template())
}

// NewFuncMap returns a template func map
func NewFuncMap() template.FuncMap {
	return template.FuncMap{
		"type":            typeFunc,
		"rawType":         rawTypeFunc,
		"zeroValue":       zeroValueFunc,
		"prependToFields": prependToFields,
	}
}

// typeFunc returns the type of the value
func typeFunc(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		return fmt.Sprintf("%T", v)
	}

	ev := reflect.New(resolveType(t)).Elem().Interface()
	return fmt.Sprintf("%T", ev)
}

// rawTypeFunc returns the raw type of the value
func rawTypeFunc(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// resolveType resolves the type of the value
func resolveType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr:
		return resolveType(t.Elem())
	}
	return t
}

// zeroValueFunc returns zero value as a string
func zeroValueFunc(v interface{}) string {
	ti := mira.NewTypeInfo(v)

	if ti.IsNillable() {
		return "nil"
	}

	if ti.IsNumeric() {
		return "0"
	}

	switch ti.T().Kind() {
	case reflect.Bool:
		return "false"
	case reflect.Struct,
		reflect.Array:
		return fmt.Sprintf("(%T{})", v)
	}

	return "\"\""

}

// prependToFields prepends a field to the list of fields
func prependToFields(field *Field, fields []*Field) []*Field {
	return append([]*Field{field}, fields...)
}
