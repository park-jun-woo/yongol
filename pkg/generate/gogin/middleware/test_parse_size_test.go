//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=size-parse
//ff:what ParseSize 테스트 — 다양한 suffix + invalid 처리

package middleware

import "testing"

func TestParseSize_ValidForms(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"1", 1},
		{"1024", 1024},
		{"1B", 1},
		{"1KiB", 1024},
		{"1MiB", 1024 * 1024},
		{"32MiB", 32 * 1024 * 1024},
		{"1GiB", 1024 * 1024 * 1024},
		{"1KB", 1000},
		{"500KB", 500 * 1000},
		{"1MB", 1000 * 1000},
		{" 2MiB ", 2 * 1024 * 1024},
		{"0", 0},
	}
	for _, c := range cases {
		got, err := ParseSize(c.in)
		if err != nil {
			t.Errorf("ParseSize(%q) err = %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("ParseSize(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestParseSize_Invalid(t *testing.T) {
	cases := []string{"", "abc", "-1MiB", "1ZZ"}
	for _, c := range cases {
		if _, err := ParseSize(c); err == nil {
			t.Errorf("ParseSize(%q) expected error, got nil", c)
		}
	}
}
