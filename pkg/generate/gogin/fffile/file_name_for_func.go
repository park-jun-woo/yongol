//ff:func feature=gen-gogin type=util control=sequence
//ff:what FileNameForFunc — PascalCase 함수명을 snake_case.go 파일명으로 변환 (1파일 1func emit 네이밍)

package fffile

import "github.com/ettle/strcase"

// FileNameForFunc converts a PascalCase Go function identifier into the
// canonical snake_case file name for a 1-file-1-func emit.
//
// Empty input yields "" so callers can detect the no-op case and skip the
// write. The ".go" suffix is always appended when a name is produced so the
// result is a valid Go source file name.
//
// Examples:
//
//	FileNameForFunc("ActivateWorkflow")   // "activate_workflow.go"
//	FileNameForFunc("convertWorkflow")    // "convert_workflow.go"
//	FileNameForFunc("HTTPServer")         // "http_server.go"
func FileNameForFunc(funcName string) string {
	if funcName == "" {
		return ""
	}
	return strcase.ToSnake(funcName) + ".go"
}
