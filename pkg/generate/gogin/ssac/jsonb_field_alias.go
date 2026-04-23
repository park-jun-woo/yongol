//ff:type feature=gen-gogin type=model
//ff:what jsonbFieldAlias — writeConvertFunc 내부 JSONB 필드의 json/api/db 이름과 로컬 변수명 집합

package ssac

// jsonbFieldAlias carries the bookkeeping for one JSONB property within
// writeConvertFunc — the OpenAPI json name, its PascalCase api-side
// field, the sqlc-side PascalCase row field, and the local variable
// name that holds the unmarshalled map before the struct literal.
type jsonbFieldAlias struct {
	jsonName string
	apiField string
	dbField  string
	localVar string
}
