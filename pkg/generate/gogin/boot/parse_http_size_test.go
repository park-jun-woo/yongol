//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what parseHTTPSize — middleware.ParseSize 래퍼: 빈 문자열/에러를 ok=false 로 통일

package boot

import "testing"

func TestParseHTTPSize(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		want   int64
		wantOk bool
	}{
		{"empty string", "", 0, false},
		{"invalid", "not-a-size", 0, false},
		{"one MiB", "1MiB", 1 << 20, true},
		{"bytes", "1024", 1024, true},
	}
	for _, c := range cases {
		got, ok := parseHTTPSize(c.in)
		if ok != c.wantOk {
			t.Errorf("%s: ok = %v, want %v", c.name, ok, c.wantOk)
		}
		if ok && got != c.want {
			t.Errorf("%s: got = %d, want %d", c.name, got, c.want)
		}
	}
}
