//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockSessionInit — session.Init (postgres infra 어댑터 또는 memory) 블록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockSessionInit_Memory(t *testing.T) {
	block := blockSessionInit(prepared.Session{Backend: "memory"}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "session.Init(session.NewMemorySession())") {
		t.Errorf("memory backend must use NewMemorySession, got:\n%s", body)
	}
}
