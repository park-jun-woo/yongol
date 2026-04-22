//ff:type feature=gen-splitter type=model
//ff:what Tool — splitter 가 인식하는 외부 코드젠 도구 (oapi-codegen / sqlc)
package splitter

// Tool identifies the external code generator whose output is being split.
// Each tool has its own filename suffix convention and list of preserved
// files that must not be split (e.g. sqlc's querier.go).
type Tool string

const (
	// ToolOAPICodegen — oapi-codegen output (*.gen.go).
	ToolOAPICodegen Tool = "oapi-codegen"
	// ToolSQLC — sqlc output (models.go, *.sql.go, querier.go, db.go).
	ToolSQLC Tool = "sqlc"
)
