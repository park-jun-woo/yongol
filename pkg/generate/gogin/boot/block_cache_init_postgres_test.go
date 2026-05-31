//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockCacheInit — cache.Init (postgres infra 어댑터 또는 memory) 블록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockCacheInit_Postgres(t *testing.T) {
	block := blockCacheInit(prepared.Cache{Backend: "postgres"}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "cache.Init(infracache.NewPostgres(queries))") {
		t.Errorf("postgres backend must wire infra adapter, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `infracache "example.com/zenflow/internal/infra/cache"`) {
		t.Errorf("postgres backend must import infra cache, got:\n%v", block.Imports)
	}
}
