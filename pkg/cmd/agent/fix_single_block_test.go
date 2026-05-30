//ff:func feature=agent type=test control=selection
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증

package agent

import (
	"bytes"
	"strings"
	"testing"
)

func TestFixSingleBlockDefaultLayer(t *testing.T) {
	// An unsupported layer returns false before any extraction.
	content := "x"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerDDL, "db/x.sql", "/tmp/db/x.sql", &content, "Op", "", "", nil) {
		t.Fatal("expected false for non-split layer")
	}
}

func TestFixSingleBlockExtractFails(t *testing.T) {
	// operationId not present → extract error → generateNewBlock, which fails to
	// find the SSaC file and returns false. Exercises the extractErr branch.
	content := "paths: {}\n"
	var out bytes.Buffer
	got := fixSingleBlock(&out, Config{}, layerOpenAPI, "api/openapi.yaml", t.TempDir()+"/api/openapi.yaml",
		&content, "MissingOp", "desc", "path", []string{"err"})
	if got {
		t.Fatal("expected false when block missing and SSaC unavailable")
	}
	if !strings.Contains(out.String(), "skipped generate") {
		t.Errorf("expected generate-skip message, got: %q", out.String())
	}
}

func TestFixSingleBlockLLMError(t *testing.T) {
	// A valid OpenAPI block is extracted, then llmCall fails (unsupported
	// backend), exercising the LLM-error branch.
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n      responses:\n        '200': {}\n"
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "none"}
	got := fixSingleBlock(&out, cfg, layerOpenAPI, "api/openapi.yaml", "/tmp/api/openapi.yaml",
		&content, "ListUsers", "desc", "path", []string{"S-01: error"})
	if got {
		t.Fatal("expected false when LLM call fails")
	}
	if !strings.Contains(out.String(), "skipped block") {
		t.Errorf("expected skipped-block message, got: %q", out.String())
	}
}

func TestFixSingleBlockRegoExtract(t *testing.T) {
	// layerRego routes through extractRegoBlock; a missing op triggers the
	// extract-error → generateNewBlock fallback.
	content := "package x\n"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerRego, "policy/x.rego", t.TempDir()+"/policy/x.rego",
		&content, "MissingOp", "d", "p", []string{"err"}) {
		t.Fatal("expected false for rego with missing op + no SSaC")
	}
}

func TestFixSingleBlockHurlExtract(t *testing.T) {
	// layerHurl routes through extractHurlBlock; a missing op triggers the
	// extract-error → generateNewBlock fallback.
	content := "GET https://x\n"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerHurl, "tests/x.hurl", t.TempDir()+"/tests/x.hurl",
		&content, "MissingOp", "d", "p", []string{"err"}) {
		t.Fatal("expected false for hurl with missing op + no SSaC")
	}
}
