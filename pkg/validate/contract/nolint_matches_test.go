//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestNolintMatches — 단일 주석이 nolint:<rule> 지시를 포함하는지 판정 검증

package contract

import "testing"

func TestNolintMatches(t *testing.T) {
	tests := []struct {
		name string
		text string
		rule string
		want bool
	}{
		{"single match", "// nolint:panic", "panic", true},
		{"no space", "//nolint:prv-14 extra", "prv-14", true},
		{"comma list first", "// nolint:prv-12,prv-17", "prv-12", true},
		{"comma list second", "// nolint:prv-12,prv-17", "prv-17", true},
		{"uppercase comment", "// NOLINT:PRV-13", "prv-13", true},
		{"rule mismatch", "// nolint:panic", "prv-10", false},
		{"no directive", "// just a comment", "panic", false},
		{"empty directive", "// nolint:", "panic", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nolintMatches(tt.text, tt.rule); got != tt.want {
				t.Fatalf("nolintMatches(%q, %q) = %v, want %v", tt.text, tt.rule, got, tt.want)
			}
		})
	}
}
