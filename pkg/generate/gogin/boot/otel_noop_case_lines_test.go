//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelNoopCaseLines — switch "noop" case 본문 (exporter 없음)

package boot

import (
	"strings"
	"testing"
)

func TestOtelNoopCaseLines(t *testing.T) {
	lines := otelNoopCaseLines()
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, `case "noop":`) {
		t.Errorf("noop branch must open `case \"noop\":`, got:\n%s", body)
	}
	if !strings.Contains(body, "spanExporter = nil") {
		t.Errorf("noop branch must nil out spanExporter, got:\n%s", body)
	}
}
