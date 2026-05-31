//ff:func feature=agent type=test control=sequence
//ff:what TestWriteOpenAPIPathContext — 테이블 관련 op들의 path 블록 기록, 파일 부재 시 무기록 검증
package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestWriteOpenAPIPathContext(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n  /orgs:\n    get:\n      operationId: ListOrgs\n"
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	lookup := map[string]features.Feature{
		"ListUsers": {Op: "ListUsers", Table: "users"},
		"ListOrgs":  {Op: "ListOrgs", Table: "orgs"},
	}

	var b strings.Builder
	writeOpenAPIPathContext(&b, dir, lookup, "users")
	out := b.String()
	if !strings.Contains(out, "OpenAPI path block (ListUsers):") {
		t.Errorf("missing ListUsers block: %q", out)
	}
	if strings.Contains(out, "ListOrgs") {
		t.Errorf("unrelated table op leaked: %q", out)
	}

	// Missing file: nothing.
	var b2 strings.Builder
	writeOpenAPIPathContext(&b2, t.TempDir(), lookup, "users")
	if b2.Len() != 0 {
		t.Errorf("missing file wrote %q", b2.String())
	}
}
