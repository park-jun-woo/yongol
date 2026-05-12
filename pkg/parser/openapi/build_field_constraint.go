//ff:func feature=manifest type=parser control=sequence
//ff:what buildFieldConstraint — 단일 property에서 FieldConstraint 생성
package openapi

import "github.com/getkin/kin-openapi/openapi3"

func buildFieldConstraint(prop *openapi3.Schema, required bool) FieldConstraint {
	var typeName string
	if types := prop.Type.Slice(); len(types) > 0 {
		typeName = types[0]
	}
	fc := FieldConstraint{
		Type:     typeName,
		Format:   prop.Format,
		Required: required,
	}
	if prop.MaxLength != nil {
		v := int(*prop.MaxLength)
		fc.MaxLength = &v
	}
	if prop.MinLength != 0 {
		v := int(prop.MinLength)
		fc.MinLength = &v
	}
	if prop.Min != nil {
		v := *prop.Min
		fc.Minimum = &v
	}
	if prop.Max != nil {
		v := *prop.Max
		fc.Maximum = &v
	}
	if prop.Pattern != "" {
		fc.Pattern = prop.Pattern
	}
	if len(prop.Enum) > 0 {
		fc.Enum = enumToStrings(prop.Enum)
	}
	return fc
}
