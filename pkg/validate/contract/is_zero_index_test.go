//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIsZeroIndex — expr 가 리터럴 0 인지 판정 검증

package contract

import "testing"

func TestIsZeroIndex(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"literal zero", "0", true},
		{"literal one", "1", false},
		{"ident i", "i", false},
		{"float zero", "0.0", false},
		{"hex zero", "0x0", false},
		{"string", "\"0\"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroIndex(mustExpr(t, tt.src)); got != tt.want {
				t.Fatalf("isZeroIndex(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
}
