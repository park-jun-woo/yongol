//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isSemVerMajor — "v2", "v3" 같은 Go 모듈 버전 suffix 판정

package boot

import "testing"

func TestIsSemVerMajor(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"v2", "v2", true},
		{"v3", "v3", true},
		{"v10", "v10", true},
		{"uppercase V2", "V2", true},
		{"v3alpha is package name", "v3alpha", false},
		{"bare v", "v", false},
		{"empty", "", false},
		{"no v prefix", "23", false},
		{"version word", "version", false},
		{"v with letters", "vx", false},
	}
	for _, c := range cases {
		if got := isSemVerMajor(c.in); got != c.want {
			t.Errorf("%s: isSemVerMajor(%q) = %v, want %v", c.name, c.in, got, c.want)
		}
	}
}
