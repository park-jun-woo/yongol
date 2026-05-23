//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what isForbiddenImport — 금지 import prefix 매칭/비매칭 검증

package funcspec

import "testing"

func TestIsForbiddenImport(t *testing.T) {
	cases := []struct {
		name string
		imp  string
		want bool
	}{
		{name: "database_sql_exact", imp: "database/sql", want: true},
		{name: "database_sql_subpkg", imp: "database/sql/driver", want: true},
		{name: "net_http_exact", imp: "net/http", want: true},
		{name: "net_http_subpkg", imp: "net/http/httptest", want: true},
		{name: "net_rpc_exact", imp: "net/rpc", want: true},
		{name: "net_rpc_subpkg", imp: "net/rpc/jsonrpc", want: true},
		{name: "grpc_exact", imp: "google.golang.org/grpc", want: true},
		{name: "grpc_subpkg", imp: "google.golang.org/grpc/codes", want: true},
		{name: "allowed_fmt", imp: "fmt", want: false},
		{name: "allowed_strings", imp: "strings", want: false},
		{name: "allowed_os", imp: "os", want: false},
		{name: "allowed_io", imp: "io", want: false},
		{name: "allowed_bufio", imp: "bufio", want: false},
		{name: "allowed_net_not_http", imp: "net", want: false},
		{name: "empty_string", imp: "", want: false},
		{name: "partial_prefix_no_slash", imp: "database/sqlfoo", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := isForbiddenImport(c.imp)
			if got != c.want {
				t.Errorf("isForbiddenImport(%q) = %v, want %v", c.imp, got, c.want)
			}
		})
	}
}
