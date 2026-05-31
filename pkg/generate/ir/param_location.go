//ff:type feature=gen-ir type=model
//ff:what ParamLocation -- HTTP 요청 필드의 출처 분류 (path/query/body/var/literal/user)

package ir

// ParamLocation classifies where a request field originates in the HTTP
// request. Mirrors the classification that gogin/ssac/methodGen performs
// via addParam (path/query) and mapRequestValue (body fallback).
type ParamLocation string

const (
	LocPath    ParamLocation = "path"    // OpenAPI path parameter
	LocQuery   ParamLocation = "query"   // OpenAPI query parameter
	LocBody    ParamLocation = "body"    // OpenAPI request body property
	LocVar     ParamLocation = "var"     // previous sequence result variable
	LocLiteral ParamLocation = "literal" // inline literal value
	LocUser    ParamLocation = "user"    // currentUser field
)
