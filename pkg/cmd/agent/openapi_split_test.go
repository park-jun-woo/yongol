//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증

package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockSuccess(t *testing.T) {
	content := strings.Join([]string{
		"paths:",
		"  /users:",
		"    get:",
		"      operationId: ListUsers",
		"      summary: list",
		"",
		"    post:",
		"      operationId: CreateUser",
		"  /orgs:",
		"    get:",
		"      operationId: ListOrgs",
	}, "\n")

	block, start, end, err := extractOpenAPIBlock(content, "ListUsers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "/users") || !strings.Contains(block, "operationId: ListUsers") {
		t.Errorf("block missing expected content: %q", block)
	}
	if strings.Contains(block, "CreateUser") {
		// The block should stop before the post method's operationId leaks... but
		// the path block spans the whole /users path (both methods). Just ensure
		// it does not run into /orgs.
		if strings.Contains(block, "/orgs") {
			t.Errorf("block should not include /orgs: %q", block)
		}
	}
	if start < 0 || end <= start {
		t.Errorf("bad line range: start=%d end=%d", start, end)
	}
}

func TestExtractOpenAPIBlockOpNotFound(t *testing.T) {
	_, _, _, err := extractOpenAPIBlock("paths:\n  /x:\n    get:\n      operationId: Other\n", "Missing")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected not-found error, got: %v", err)
	}
}

func TestExtractOpenAPIBlockMethodNotFound(t *testing.T) {
	// operationId with no HTTP method line above it.
	content := "info:\n  operationId: Weird\n"
	_, _, _, err := extractOpenAPIBlock(content, "Weird")
	if err == nil || !strings.Contains(err.Error(), "method line") {
		t.Fatalf("expected method-line error, got: %v", err)
	}
}

func TestExtractOpenAPIBlockPathNotFound(t *testing.T) {
	// HTTP method at column 0 (indent 0): no line with smaller indent exists, so
	// the path line cannot be found.
	content := "get:\n  operationId: Flat\n"
	_, _, _, err := extractOpenAPIBlock(content, "Flat")
	if err == nil || !strings.Contains(err.Error(), "path line") {
		t.Fatalf("expected path-line error, got: %v", err)
	}
}

func TestExtractOpenAPIBlockBlankBeforeMethod(t *testing.T) {
	// A blank line between the path key and the HTTP method exercises the
	// empty-line continue inside the path-line search loop.
	content := strings.Join([]string{
		"paths:",
		"  /users:",
		"",
		"    get:",
		"      operationId: ListUsers",
	}, "\n")
	block, _, _, err := extractOpenAPIBlock(content, "ListUsers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "/users") {
		t.Errorf("block missing path: %q", block)
	}
}
