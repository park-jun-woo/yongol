//ff:func feature=gen-gogin type=test control=sequence topic=health
//ff:what TestBlockHealth_NoDDLSkipsHelper — DDL 없으면 /ready 헬퍼/임포트 생략

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockHealth_NoDDLSkipsHelper — without DDL the /ready endpoint is
// a static 200 and no helper/import is emitted.
func TestBlockHealth_NoDDLSkipsHelper(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
	}
	block := blockHealth(fs)
	funcs := strings.Join(block.Funcs, "\n")
	imports := strings.Join(block.Imports, "\n")

	if funcs != "" {
		t.Fatalf("no DDL: must not emit readyHandlerWithDB helper, got:\n%s", funcs)
	}
	if strings.Contains(imports, `pgxpool`) {
		t.Fatalf("no DDL: must not import pgxpool, got:\n%s", imports)
	}
}
