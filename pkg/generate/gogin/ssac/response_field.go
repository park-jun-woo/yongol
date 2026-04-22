//ff:type feature=gen-gogin type=model
//ff:what responseField — OpenAPI 200 응답 스키마의 단일 필드 기술

package ssac

// responseField describes one field in the 200 response body.
type responseField struct {
	JSONName   string // "workflow"
	GoName     string // "Workflow"
	RefType    string // "Workflow" ($ref schema name) or ""
	IsArray    bool
	IsRequired bool // OpenAPI required[] membership — non-pointer in generated Go
}
