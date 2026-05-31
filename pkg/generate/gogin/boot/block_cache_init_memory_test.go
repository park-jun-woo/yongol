//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockCacheInit — cache.Init (postgres infra 어댑터 또는 memory) 블록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockCacheInit_Memory(t *testing.T) {
	block := blockCacheInit(prepared.Cache{Backend: "memory"}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "cache.Init(cache.NewMemoryCache())") {
		t.Errorf("memory backend must use NewMemoryCache, got:\n%s", body)
	}
	if strings.Contains(strings.Join(block.Imports, "\n"), "infracache") {
		t.Errorf("memory backend must not import infra adapter, got:\n%v", block.Imports)
	}
}
