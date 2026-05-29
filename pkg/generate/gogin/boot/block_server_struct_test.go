//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockServerStruct — Server struct 초기화 블록

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockServerStruct(t *testing.T) {
	block := blockServerStruct(&yongol.Fullstack{}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"srv := &service.Server{",
		"DB: pool,",
		"Queries: queries,",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("blockServerStruct missing %q, got:\n%s", must, body)
		}
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"example.com/zenflow/internal/service"`) {
		t.Errorf("must import service, got:\n%v", block.Imports)
	}
}
