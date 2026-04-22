//ff:func feature=gen-splitter type=util control=selection
//ff:what suffixFor — 도구 + 원본 파일 역할에 따라 분할 파일 suffix 결정
package splitter

import "go/ast"

// suffixFor picks the split file suffix. oapi-codegen always uses .gen.go.
// For sqlc, struct/type specs from models.go get .model.go so that the
// sqlc row-type convention stays visible; everything else (funcs, consts,
// non-models type aliases) lands on .sql.go.
func suffixFor(tool Tool, isModelsFile bool, decl ast.Decl) string {
	switch tool {
	case ToolOAPICodegen:
		return ".gen.go"
	case ToolSQLC:
		return sqlcSuffix(isModelsFile, decl)
	default:
		return ".go"
	}
}
