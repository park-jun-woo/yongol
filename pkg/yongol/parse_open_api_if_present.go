//ff:func feature=orchestrator type=loader control=sequence
//ff:what OpenAPI 탐지 시 parseOpenAPI 로 doc/lines 로드 후 request/response 제약 추출
package yongol

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// parseOpenAPIIfPresent loads the single-site OpenAPI document via parseOpenAPI,
// then extracts request/response constraints. Load failure is collected as a
// diagnostic (doc stays nil); line-index failure only degrades line reporting.
// The ExtractRequest/ResponseConstraints calls MUST remain here — 19 readers
// depend on fs.RequestConstraints / fs.ResponseConstraints.
func parseOpenAPIIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindOpenAPI]
	if !ok {
		return
	}
	doc, lines, diags := parseOpenAPI(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if doc == nil {
		return
	}
	fs.OpenAPIDoc = doc
	fs.OpenAPILines = lines
	fs.RequestConstraints = oapiparser.ExtractRequestConstraints(doc, lines)
	fs.ResponseConstraints = oapiparser.ExtractResponseConstraints(doc, lines)
}
