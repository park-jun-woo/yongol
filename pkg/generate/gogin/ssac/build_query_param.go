//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildQueryParam — ParameterRef → queryParam 메타데이터 (format/enum/nullable)

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// buildQueryParam collects the metadata yongol needs to generate correctly
// typed Go access code for one query parameter: primitive GoType, enum alias
// name (oapi-codegen naming: "<OperationId>Params<PascalCase(name)>"),
// required/nullable flags and format.
func buildQueryParam(p *openapi3.ParameterRef, operationId string) queryParam {
	qp := queryParam{GoType: "string"}
	if p == nil || p.Value == nil {
		return qp
	}
	qp.IsRequired = p.Value.Required
	if p.Value.Schema == nil || p.Value.Schema.Value == nil {
		return qp
	}
	schema := p.Value.Schema.Value
	qp.Format = schema.Format
	qp.IsNullable = schema.Nullable
	if types := schema.Type; types != nil && len(*types) > 0 {
		t := (*types)[0]
		if t == "integer" && schema.Format == "int64" {
			qp.GoType = "integer64"
		} else {
			qp.GoType = t
		}
	}
	if len(schema.Enum) > 0 {
		qp.IsEnum = true
		qp.EnumTypeName = pascalCase(operationId) + "Params" + pascalCase(p.Value.Name)
	}
	return qp
}
