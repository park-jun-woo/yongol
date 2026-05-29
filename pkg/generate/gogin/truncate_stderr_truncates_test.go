//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTruncateStderrTruncates — limit 초과 시 자르고 마커 추가

package gogin

import (
	"strings"
	"testing"
)

// TestTruncateStderrTruncates — string over the limit is cut and marked.
func TestTruncateStderrTruncates(t *testing.T) {
	in := strings.Repeat("x", 50)
	got := truncateStderr(in, 10)
	want := strings.Repeat("x", 10) + "...(truncated)"
	if got != want {
		t.Fatalf("truncateStderr = %q, want %q", got, want)
	}
}
