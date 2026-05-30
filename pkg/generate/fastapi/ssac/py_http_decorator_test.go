//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestPyHTTPDecorator — HTTP 메서드 → FastAPI 데코레이터 이름 매핑

package ssac

import "testing"

func TestPyHTTPDecorator(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"GET", "get"},
		{"post", "post"},
		{"Put", "put"},
		{"DELETE", "delete"},
		{"patch", "patch"},
		{"UNKNOWN", "get"},
		{"", "get"},
	}
	for _, c := range cases {
		if got := pyHTTPDecorator(c.in); got != c.want {
			t.Errorf("pyHTTPDecorator(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
