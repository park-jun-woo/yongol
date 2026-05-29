//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestCallClosesVar — CallExpr 가 varName 의 .Close() 호출인지 판정 검증

package contract

import "testing"

func TestCallClosesVar(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		varName string
		want    bool
	}{
		{"direct close", "f.Close()", "f", true},
		{"chained close", "resp.Body.Close()", "resp", true},
		{"wrong var", "f.Close()", "g", false},
		{"non close method", "f.Read(b)", "f", false},
		{"free function", "Close()", "f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callClosesVar(mustCall(t, tt.src), tt.varName); got != tt.want {
				t.Fatalf("callClosesVar(%q, %q) = %v, want %v", tt.src, tt.varName, got, tt.want)
			}
		})
	}
	if callClosesVar(nil, "f") {
		t.Fatal("nil call should return false")
	}
}
