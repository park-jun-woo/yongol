package openapi

// httpMethods OpenAPI 가 PathItem 안에서 operation 으로 인식하는 HTTP 메서드들.
var httpMethods = map[string]bool{
	"get":     true,
	"put":     true,
	"post":    true,
	"delete":  true,
	"options": true,
	"head":    true,
	"patch":   true,
	"trace":   true,
}
