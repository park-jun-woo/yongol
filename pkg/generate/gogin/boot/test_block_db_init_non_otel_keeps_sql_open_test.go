//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockDBInit_NonOtelKeepsSqlOpen — tracing 미활성 시 sql.Open 유지 회귀

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockDBInit_NonOtelKeepsSqlOpen(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{},
		},
	}
	block := blockDBInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `sql.Open("postgres"`) {
		t.Fatalf("non-tracing db-init must keep sql.Open, got:\n%s", body)
	}
	if strings.Contains(body, "otelsql") {
		t.Fatalf("non-tracing db-init must NOT import otelsql, got:\n%s", body)
	}
}
