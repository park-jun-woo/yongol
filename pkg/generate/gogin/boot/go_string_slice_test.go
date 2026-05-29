//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what goStringSlice — []string을 Go []string{...} 리터럴로 렌더링

package boot

import "testing"

func TestGoStringSlice(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want string
	}{
		{"empty", nil, `[]string{}`},
		{"empty slice", []string{}, `[]string{}`},
		{"single", []string{"GET"}, `[]string{"GET"}`},
		{"multiple", []string{"GET", "POST"}, `[]string{"GET", "POST"}`},
		{"quotes escaped", []string{`a"b`}, `[]string{"a\"b"}`},
	}
	for _, c := range cases {
		if got := goStringSlice(c.in); got != c.want {
			t.Errorf("%s: goStringSlice(%v) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
