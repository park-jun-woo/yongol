//ff:func feature=validate-contract type=test control=selection topic=preserve-safety
//ff:what TestIsScanCall — call 이 row/rows/r.Scan(...) 호출인지 판정 검증

package contract

import "testing"

func TestIsScanCall(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"row scan", "row.Scan(&x)", true},
		{"rows scan", "rows.Scan(&x)", true},
		{"r scan", "r.Scan(&x)", true},
		{"suffix row", "userRow.Scan(&x)", true},
		{"no args", "row.Scan()", false},
		{"non scan method", "row.Next()", false},
		{"unrelated receiver", "scanner.Scan(&x)", false},
		{"free function", "Scan(&x)", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScanCall(mustCall(t, tt.src)); got != tt.want {
				t.Fatalf("isScanCall(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
	if isScanCall(nil) {
		t.Fatal("nil call should return false")
	}
}
