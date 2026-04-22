//ff:func feature=orchestrator type=loader control=sequence
//ff:what OpenAPI 탐지 시 kin-openapi 로드 + yaml.v3 라인 인덱스 + request/response 제약 추출
package yongol

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// parseOpenAPIIfPresent loads the OpenAPI document, builds a line index and
// extracts request/response constraints. Load failure is collected as a
// diagnostic; line-index failure only degrades line reporting.
func parseOpenAPIIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindOpenAPI]
	if !ok {
		return
	}
	doc, err := openapi3.NewLoader().LoadFromFile(d.Path)
	if err != nil {
		fs.ParseDiagnostics = append(fs.ParseDiagnostics, diagnostic.Diagnostic{
			File:    d.Path,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "OpenAPI load error: " + err.Error(),
		})
		return
	}
	fs.OpenAPIDoc = doc
	// Raw-parse the same file once more with yaml.v3 to build a line-number
	// index. On failure (e.g. file permissions) keep Fullstack alive — only
	// line information is missing; kin-openapi results are intact so
	// validation can still proceed.
	lines, lerr := oapiparser.BuildLineIndex(d.Path)
	if lerr != nil {
		fs.ParseDiagnostics = append(fs.ParseDiagnostics, diagnostic.Diagnostic{
			File:    d.Path,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelWarning,
			Message: "OpenAPI line index build error: " + lerr.Error(),
		})
	}
	fs.OpenAPILines = lines
	fs.RequestConstraints = oapiparser.ExtractRequestConstraints(doc, lines)
	fs.ResponseConstraints = oapiparser.ExtractResponseConstraints(doc, lines)
}
