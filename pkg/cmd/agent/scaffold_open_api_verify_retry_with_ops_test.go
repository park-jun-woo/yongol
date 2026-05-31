//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldOpenAPIVerifyRetry — verify 성공 / 미귀속 verify 에러 stop / 귀속 op 재시도+재조립 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldOpenAPIVerifyRetryWithOps(t *testing.T) {
	// An invalid document whose error references a line within an op's offset
	// range drives the retry loop and the reassembly. retryFailedOp fails for the
	// unsupported backend but the loop still executes.
	doc := "openapi: \"3.0.0\"\ninfo:\n  title: T\n  version: x\npaths:\n  /ping:\n    get: bad\n"
	offsets := []pathOffset{{Op: "Ping", Path: "/ping", StartLine: 1, EndLine: 100}}
	featByOp := map[string]features.Feature{"Ping": {Op: "Ping", Path: "/ping", Table: "t"}}
	pathBlocks := map[string]any{"/ping": map[string]any{"get": "x"}}
	pathToOps := map[string][]string{"/ping": {"Ping"}}
	opToPath := map[string]string{"Ping": "/ping"}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m", SpecsDir: t.TempDir()}
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Ping", Path: "/ping"}}}

	res := scaffoldOpenAPIVerifyRetry(&doc, &offsets, pathBlocks, pathToOps, opToPath, featByOp,
		nil, false, 0, cfg, &out, ff)
	if res.verified || res.stopped {
		t.Errorf("expected a plain retry result (not verified/stopped), got %+v", res)
	}
}
