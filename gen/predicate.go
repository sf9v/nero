package gen

import (
	"bytes"
	"text/template"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/comparison"
)

func newPredicateFile(schema *nero.Schema) (*bytes.Buffer, error) {
	tmpl, err := template.New("predicates.tmpl").
		Funcs(nero.NewFuncMap()).Parse(predicatesTmpl)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	data := struct {
		EqOps   []comparison.Operator
		LtGtOps []comparison.Operator
		NullOps []comparison.Operator
		InOps   []comparison.Operator
		Schema  *nero.Schema
	}{
		EqOps: []comparison.Operator{
			comparison.Eq,
			comparison.NotEq,
		},
		LtGtOps: []comparison.Operator{
			comparison.Gt,
			comparison.GtOrEq,
			comparison.Lt,
			comparison.LtOrEq,
		},
		NullOps: []comparison.Operator{
			comparison.IsNull,
			comparison.IsNotNull,
		},
		InOps: []comparison.Operator{
			comparison.In,
			comparison.NotIn,
		},
		Schema: schema,
	}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

const predicatesTmpl = `
// Code generated by nero, DO NOT EDIT.
package {{.Schema.PkgName}}

import (
	"github.com/lib/pq"
	"github.com/sf9v/nero/comparison"
	{{range $import := .Schema.Imports -}}
		"{{$import}}"
	{{end -}}
)

{{ $cols := prependToColumns .Schema.Identity .Schema.Columns }}

{{range $col := $cols -}}
	{{if $col.CanHavePreds -}}
        {{ range $op := $.EqOps }} 
            // {{$col.FieldName}}{{$op.String}} applies "{{$op.Desc}}" operator on "{{$col.Name}}" column
            func {{$col.FieldName}}{{$op.String}} ({{$col.Identifier}} {{printf "%T" $col.TypeInfo.V}}) comparison.PredFunc {
                return func(preds []*comparison.Predicate) []*comparison.Predicate {
                    return append(preds, &comparison.Predicate{
                        Col: "{{$col.Name}}",
                        Op: comparison.{{$op.String}},
                        {{if and ($col.IsArray) (ne $col.IsValueScanner true) -}}
                            Arg: pq.Array({{$col.Identifier}}),
                        {{else -}}
                            Arg: {{$col.Identifier}},
                        {{end -}}
                    })
                }
            }

            {{if $col.IsComparable -}}
                // {{$col.FieldName}}{{$op.String}} applies "{{$op.Desc}}" operator on "{{$col.Name}}" column
                func {{$col.FieldName}}{{$op.String}}Col (col Column) comparison.PredFunc {
                    return func(preds []*comparison.Predicate) []*comparison.Predicate {
                        return append(preds, &comparison.Predicate{
                            Col: "{{$col.Name}}",
                            Op: comparison.{{$op.String}},
                            Arg: col,
                        })
                    }
                }
            {{end}}
        {{end}}

        {{ range $op := $.LtGtOps }}
            {{if $col.TypeInfo.IsNumeric }}
                // {{$col.FieldName}}{{$op.String}} applies "{{$op.Desc}}" operator on "{{$col.Name}}" column
                func {{$col.FieldName}}{{$op.String}} ({{$col.Identifier}} {{printf "%T" $col.TypeInfo.V}}) comparison.PredFunc {
                    return func(preds []*comparison.Predicate) []*comparison.Predicate {
                        return append(preds, &comparison.Predicate{
                            Col: "{{$col.Name}}",
                            Op: comparison.{{$op.String}},
                            {{if and ($col.IsArray) (ne $col.IsValueScanner true) -}}
                                Arg: pq.Array({{$col.Identifier}}),
                            {{else -}}
                                Arg: {{$col.Identifier}},
                            {{end -}}
                        })
                    }
                }
            {{end}}
        {{end }}

        {{ range $op := $.NullOps }}
            {{if $col.IsNillable}}
                // {{$col.FieldName}}{{$op.String}} applies "{{$op.Desc}}" operator on "{{$col.Name}}" column
                func {{$col.FieldName}}{{$op.String}} () comparison.PredFunc {
                    return func(preds []*comparison.Predicate) []*comparison.Predicate {
                        return append(preds, &comparison.Predicate{
                            Col: "{{$col.Name}}",
                            Op: comparison.{{$op.String}},
                        })
                    }
                }
            {{end}}
        {{end}}

        {{ range $op := $.InOps }}
            // {{$col.FieldName}}{{$op.String}} applies "{{$op.Desc}}" operator on "{{$col.Name}}" column
            func {{$col.FieldName}}{{$op.String}} ({{$col.IdentifierPlural}} {{printf "...%T" $col.TypeInfo.V}}) comparison.PredFunc {
                args := []interface{}{}
                for _, v := range {{$col.IdentifierPlural}} {
                    args = append(args, v)
                }

                return func(preds []*comparison.Predicate) []*comparison.Predicate {
                    return append(preds, &comparison.Predicate{
                        Col: "{{$col.Name}}",
                        Op: comparison.{{$op.String}},
                        Arg: args,
                    })
                }
            }
        {{end}}
	{{end}}
{{end -}}
`