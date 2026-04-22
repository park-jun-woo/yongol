//ff:func feature=gen-splitter type=util control=selection
//ff:what isPreservedFile — 도구별 분할 금지 파일 여부 판정
package splitter

import "path/filepath"

// isPreservedFile reports whether name is a file that the splitter must
// leave untouched. sqlc's querier.go and db.go are the interface core and
// DBTX/New constructor respectively — splitting them breaks the runtime
// coupling that sqlc assumes. Other tools may add preserved files here.
func isPreservedFile(name string, tool Tool) bool {
	base := filepath.Base(name)
	switch tool {
	case ToolSQLC:
		switch base {
		case "querier.go", "db.go":
			return true
		}
		return false
	case ToolOAPICodegen:
		return false
	default:
		return false
	}
}
