//ff:type feature=validate type=model topic=ssac-openapi
//ff:what errStatusTypes — SSaC 시퀀스 타입별 기본 ErrStatus 매핑

package openapi_ssac

// errStatusTypes maps SSaC sequence types that produce error responses to
// their default HTTP status codes.
var errStatusTypes = map[string]int{
	"empty":  404,
	"exists": 409,
	"state":  409,
	"auth":   403,
	"call":   500,
}
