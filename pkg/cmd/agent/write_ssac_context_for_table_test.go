//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestWriteSSaCContextForTable — 테이블 관련 op들의 SSaC 컨텍스트 기록, 무관련 op 제외 검증

package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestWriteSSaCContextForTable(t *testing.T) {
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svc, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svc, "Login.ssac"), []byte("func Login() {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svc, "ListOrgs.ssac"), []byte("func ListOrgs() {}"), 0o644); err != nil {
		t.Fatal(err)
	}

	lookup := map[string]features.Feature{
		"Login":    {Op: "Login", Table: "users"},
		"ListOrgs": {Op: "ListOrgs", Table: "orgs"},
	}

	var b strings.Builder
	writeSSaCContextForTable(&b, dir, lookup, "users")
	out := b.String()
	if !strings.Contains(out, "SSaC (Login.ssac):") {
		t.Errorf("missing Login ssac: %q", out)
	}
	if strings.Contains(out, "ListOrgs") {
		t.Errorf("unrelated table op leaked: %q", out)
	}
}
