//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=size-parse
//ff:what TestParseSize_ValidForms — 다양한 단위 suffix 파싱 성공 확인

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
