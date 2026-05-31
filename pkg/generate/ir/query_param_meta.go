//ff:type feature=gen-ir type=model
//ff:what QueryParamMeta -- 단일 OpenAPI query 파라미터 메타데이터

package ir

// QueryParamMeta holds metadata for a single OpenAPI query parameter.
type QueryParamMeta struct {
	// Name is the OpenAPI parameter name.
	Name string
	// Type is the schema type (e.g. "string", "integer").
	Type string
	// Required is true when the parameter is required.
	Required bool
}
