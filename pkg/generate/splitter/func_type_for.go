//ff:func feature=gen-splitter type=util control=selection
//ff:what funcTypeFor — Tool 별로 arts/backend codebook.yaml 의 type= 값 선택
package splitter

// funcTypeFor maps a Tool to the codebook type= label assigned to each
// split func: oapi-codegen wrappers are HTTP handlers, sqlc methods are
// database queries. The labels match the entries in the auto-generated
// arts/backend/codebook.yaml so filefunc A2 accepts them without manual
// codebook edits downstream.
func funcTypeFor(tool Tool) string {
	switch tool {
	case ToolOAPICodegen:
		return "handler"
	case ToolSQLC:
		return "query"
	default:
		return "util"
	}
}
