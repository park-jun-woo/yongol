//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what buildCallSpanOpenLines 단위 테스트 (@call span.Start 라인 방출)

package ssac

import "testing"

func TestBuildCallSpanOpenLines(t *testing.T) {
	lines := buildCallSpanOpenLines("dashboard", "Summarize")
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d: %v", len(lines), lines)
	}
	want := `callCtx, callSpan := otel.Tracer("ssac").Start(ctx, "call.dashboard.Summarize")`
	if lines[0] != want {
		t.Errorf("line = %q, want %q", lines[0], want)
	}
}
