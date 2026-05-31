//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼
package splitter

import (
	"testing"
)

func TestSnake(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"CamelCase", "camel_case"},
		{"HTTPServer", "http_server"},
		{"APIKey", "api_key"},
		{"simple", "simple"},
		{"ID", "id"},
		{"UserID", "user_id"},
		{"", ""},
	}
	for _, c := range cases {
		if got := snake(c.in); got != c.want {
			t.Errorf("snake(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
