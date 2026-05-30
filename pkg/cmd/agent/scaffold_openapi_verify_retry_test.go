//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestScaffoldOpenAPIVerifyRetry — verify 성공 / 미귀속 verify 에러 stop / 귀속 op 재시도+재조립 분기 검증

package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

const validOpenAPIDoc = `openapi: "3.0.0"
info:
  title: T
  version: "1.0.0"
paths:
  /ping:
    get:
      operationId: Ping
      responses:
        '200':
          description: ok
`

func TestScaffoldOpenAPIVerifyRetryVerified(t *testing.T) {
	doc := validOpenAPIDoc
	var offsets []pathOffset
	var out bytes.Buffer
	res := scaffoldOpenAPIVerifyRetry(&doc, &offsets, map[string]any{}, map[string][]string{},
		map[string]string{}, map[string]features.Feature{}, nil, false, 0, Config{}, &out,
		&features.FeaturesFile{})
	if !res.verified {
		t.Errorf("expected verified=true for a valid document, got %+v; out=%q", res, out.String())
	}
}

func TestScaffoldOpenAPIVerifyRetryNoAttribution(t *testing.T) {
	// An invalid document with no offsets → no ops can be attributed → stopped.
	doc := "not: a valid openapi doc\n"
	var offsets []pathOffset
	var out bytes.Buffer
	res := scaffoldOpenAPIVerifyRetry(&doc, &offsets, map[string]any{}, map[string][]string{},
		map[string]string{}, map[string]features.Feature{}, nil, false, 0, Config{}, &out,
		&features.FeaturesFile{})
	if !res.stopped {
		t.Errorf("expected stopped=true when no ops attributable, got %+v", res)
	}
}

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
