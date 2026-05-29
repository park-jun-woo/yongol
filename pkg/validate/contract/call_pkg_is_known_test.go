//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestCallPkgIsKnown — pkg.Func 의 pkg 접두사가 알려진 패키지 집합에 속하는지 검증

package contract

import "testing"

func TestCallPkgIsKnown(t *testing.T) {
	known := map[string]bool{"billing": true, "users": true}
	tests := []struct {
		name string
		call string
		want bool
	}{
		{"known pkg", "billing.Charge", true},
		{"another known", "users.Find", true},
		{"unknown pkg", "fmt.Sprintf", false},
		{"no dot", "Charge", false},
		{"leading dot", ".Charge", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callPkgIsKnown(tt.call, known); got != tt.want {
				t.Fatalf("callPkgIsKnown(%q) = %v, want %v", tt.call, got, tt.want)
			}
		})
	}
}
