//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestLegacyColumnlessExpr — column 없는 legacy source → TS 식별자 매핑 검증

package ssac

import "testing"

func TestLegacyColumnlessExpr(t *testing.T) {
	tests := []struct {
		source string
		want   string
	}{
		{"request", "params"},
		{"currentUser", "user"},
		{"order", "order"}, // default: passthrough
	}
	for _, tt := range tests {
		if got := legacyColumnlessExpr(tt.source); got != tt.want {
			t.Errorf("legacyColumnlessExpr(%q) = %q, want %q", tt.source, got, tt.want)
		}
	}
}
