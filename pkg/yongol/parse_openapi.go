//ff:func feature=orchestrator type=loader control=sequence
//ff:what 경로의 OpenAPI 파일을 kin-openapi 로 로드하고 yaml.v3 라인 인덱스를 구축 — 단일/도메인 공용, 진단 수집
package yongol

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// parseOpenAPI loads the OpenAPI document at path with kin-openapi and builds a
// yaml.v3 line-number index for it. A load failure returns a nil doc plus an
// error diagnostic; a line-index failure is non-fatal (warning diagnostic, the
// doc and a partial index are still returned). Shared by the single-site loader
// (parseOpenAPIIfPresent) and the per-domain loop (parseDomainsIfPresent) so both
// paths build the doc/lines pair identically.
func parseOpenAPI(path string) (*openapi3.T, *oapiparser.LineIndex, []diagnostic.Diagnostic) {
	var diags []diagnostic.Diagnostic
	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "OpenAPI load error: " + err.Error(),
		})
		return nil, nil, diags
	}
	// Raw-parse the same file once more with yaml.v3 to build a line-number
	// index. On failure keep the doc alive — only line information is missing;
	// kin-openapi results are intact so validation can still proceed.
	lines, lerr := oapiparser.BuildLineIndex(path)
	if lerr != nil {
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelWarning,
			Message: "OpenAPI line index build error: " + lerr.Error(),
		})
	}
	return doc, lines, diags
}
