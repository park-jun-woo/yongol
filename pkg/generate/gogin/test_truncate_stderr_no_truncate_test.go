//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTruncateStderrNoTruncate — limit 이내면 trailing whitespace 만 제거

package gogin

import "testing"

// TestTruncateStderrNoTruncate — string fits under the limit, returned as-is
// except for trailing whitespace trimming.
func TestTruncateStderrNoTruncate(t *testing.T) {
	got := truncateStderr("hello\n", 64)
	if got != "hello" {
		t.Fatalf("truncateStderr = %q, want %q", got, "hello")
	}
}
