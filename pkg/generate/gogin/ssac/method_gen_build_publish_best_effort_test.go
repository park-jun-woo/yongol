//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what buildPublishBestEffort 단위 테스트 (queue.Publish + slog.Error best-effort 블록 + slog import)

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildPublishBestEffort(t *testing.T) {
	g := &methodGen{FuncName: "CompleteOrder"}
	seq := ssacparser.Sequence{Topic: "order.completed"}
	fields := []string{`"order_id": orderID,`}
	var imports []string

	lines := g.buildPublishBestEffort(seq, fields, &imports)
	joined := strings.Join(lines, "\n")

	if !strings.Contains(joined, `queue.Publish(ctx, "order.completed", map[string]any{`) {
		t.Errorf("missing queue.Publish opener:\n%s", joined)
	}
	if !strings.Contains(joined, `"order_id": orderID,`) {
		t.Errorf("payload fields not inlined:\n%s", joined)
	}
	if !strings.Contains(joined, `slog.Error("publish failed", "op", "CompleteOrder", "topic", "order.completed", "err", err)`) {
		t.Errorf("missing slog.Error best-effort log:\n%s", joined)
	}
	found := false
	for _, im := range imports {
		if im == `"log/slog"` {
			found = true
		}
	}
	if !found {
		t.Errorf("expected log/slog import appended, got %v", imports)
	}
}
