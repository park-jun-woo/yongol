//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockQueueInit — queue.Init + SetBackend + Subscribe + Start + defer Close
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockQueueInit_MemoryNoSubscribe(t *testing.T) {
	block := blockQueueInit(prepared.Queue{Backend: "memory"}, nil, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `queue.Init(ctx, "memory")`) {
		t.Errorf("must init memory queue, got:\n%s", body)
	}
	if strings.Contains(body, "SetBackend") {
		t.Errorf("memory backend must not call SetBackend, got:\n%s", body)
	}
	if !strings.Contains(body, "defer queue.Close()") {
		t.Errorf("must defer queue.Close(), got:\n%s", body)
	}
}
