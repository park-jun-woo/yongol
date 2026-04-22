//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what findOperationIDForFile — arts Go 파일 경로 → 추정 operationId (snake_case basename → PascalCase)

package contract

import (
	"path/filepath"
	"strings"
)

// findOperationIDForFile derives the operationId that a preserved service
// file is expected to implement from its basename. arts backends follow
// the `pkg/generate/gogin/ssac` convention of emitting one file per SSaC
// func with a snake_case filename that mirrors the PascalCase operationId.
//
// Examples:
//   - arts/backend/internal/service/activate_workflow.go → "ActivateWorkflow"
//   - arts/backend/internal/service/list_templates.go    → "ListTemplates"
//
// Returns "" when the basename cannot be interpreted (empty after trim).
func findOperationIDForFile(path string) string {
	base := filepath.Base(path)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	if name == "" {
		return ""
	}
	parts := strings.Split(name, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return b.String()
}
