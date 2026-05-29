//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestDF_06_QueueInit_HasDeferClose — blockQueueInit 템플릿이 defer queue.Close() + err 가드를 포함하는지 회귀 방지

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// TestDF_06_QueueInit_HasDeferClose asserts the queue-init template keeps
// both `defer queue.Close()` (DF-06) and the queue.Init err guard
// (DF family). A missing defer would leak the broker connection on
// shutdown; a missing guard would hide init failures.
func TestDF_06_QueueInit_HasDeferClose(t *testing.T) {
	block := blockQueueInit(prepared.Queue{Backend: "postgres"}, nil, "example.com/zenflow")
	lines := strings.Join(block.Lines, "\n")
	if !strings.Contains(lines, "defer queue.Close()") {
		t.Fatalf("queue-init must defer queue.Close() (DF-06), got:\n%s", lines)
	}
	if !strings.Contains(lines, "queue.Init(ctx,") {
		t.Fatalf("queue-init must call queue.Init, got:\n%s", lines)
	}
	if !strings.Contains(lines, "slog.Error(\"queue init\"") {
		t.Fatalf("queue-init must guard queue.Init error, got:\n%s", lines)
	}
}
