package nero

import (
	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"
)

// Schema is a schema used for generating the repository
type Schema struct {
	// pkgName is the package name of the generated files
	pkgName string
	// Collection is the name of the database collection
	collection string
	// typeInfo is the type info of the schema model
	typeInfo *mira.TypeInfo
	// identity is the identity field
	identity *Field
	// fields is the list of fields
	fields []*Field
	// imports are list of package imports
	imports []string
	// Templates is the list of custom repository templates
	templates []Template
}

// PkgName returns the pkg name
func (s *Schema) PkgName() string {
	return s.pkgName
}

// Collection returns the collection
func (s *Schema) Collection() string {
	return s.collection
}

// Identity returns the identity field
func (s *Schema) Identity() *Field {
	return s.identity
}

// Fields returns the fields
func (s *Schema) Fields() []*Field {
	return s.fields[:]
}

// Imports returns the pkg imports
func (s *Schema) Imports() []string {
	return s.imports[:]
}

// Templates returns the templates
func (s *Schema) Templates() []Template {
	return s.templates[:]
}

// TypeInfo returns the type info
func (s *Schema) TypeInfo() *mira.TypeInfo {
	return s.typeInfo
}

// TypeName returns the type name
func (s *Schema) TypeName() string {
	return s.typeInfo.Name()
}

// TypeNamePlural returns the plural form of the type name
func (s *Schema) TypeNamePlural() string {
	return inflection.Plural(s.TypeName())
}

// TypeIdentifier returns the type identifier
func (s *Schema) TypeIdentifier() string {
	return stringsx.ToLowerCamel(s.TypeName())
}

// TypeIdentifierPlural returns the plural form of type identifier
func (s *Schema) TypeIdentifierPlural() string {
	return inflection.Plural(s.TypeIdentifier())
}
