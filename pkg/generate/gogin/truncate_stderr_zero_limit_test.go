//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTruncateStderrZeroLimit — limit<=0 시 자르지 않고 trim 만 수행

package gogin

import "testing"

// TestTruncateStderrZeroLimit — limit<=0 short-circuits and just trims.
func TestTruncateStderrZeroLimit(t *testing.T) {
	got := truncateStderr("abc  \n", 0)
	if got != "abc" {
		t.Fatalf("truncateStderr(limit=0) = %q, want %q", got, "abc")
	}
}
