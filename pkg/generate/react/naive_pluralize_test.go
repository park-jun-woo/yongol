//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what naivePluralize 복수형 변환 테이블 검증

package react

import "testing"

func TestNaivePluralize(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"workflow", "workflows"},
		{"bus", "buses"},
		{"match", "matches"},
		{"box", "boxes"},
		{"buzz", "buzzes"},
		{"dish", "dishes"},
		{"template", "templates"},
	}
	for _, tt := range tests {
		got := naivePluralize(tt.in)
		if got != tt.want {
			t.Errorf("naivePluralize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
