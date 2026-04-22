//ff:func feature=gen-splitter type=util control=selection
//ff:what matchesOriginal — 파일 이름이 해당 도구의 "원본" 패턴(분할 전)에 해당하는지 판정
package splitter

import "strings"

// matchesOriginal reports whether name looks like an un-split generator
// output for tool. For sqlc we delete the source files (models.go and
// *.sql.go) unconditionally on regeneration and rely on keep[] to skip
// fresh splits. For oapi-codegen we target anything matching *.gen.go
// so stale splits from earlier runs are also cleaned up.
func matchesOriginal(name string, tool Tool) bool {
	switch tool {
	case ToolOAPICodegen:
		return strings.HasSuffix(name, ".gen.go")
	case ToolSQLC:
		if name == "models.go" {
			return true
		}
		return strings.HasSuffix(name, ".sql.go") || strings.HasSuffix(name, ".model.go")
	default:
		return false
	}
}
