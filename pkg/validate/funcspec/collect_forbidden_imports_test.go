//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what collectForbiddenImports — 금지 import 존재/부재에 따른 XFF-41 진단 생성 검증

package funcspec

import (
	"testing"
)

func TestCollectForbiddenImports(t *testing.T) {
	cases := []TestCollectForbiddenImportsCase{
		{
			name:      "no_imports_returns_empty",
			pkg:       "billing",
			funcName:  "CreateInvoice",
			line:      10,
			imports:   nil,
			wantCount: 0,
		},
		{
			name:      "allowed_imports_returns_empty",
			pkg:       "billing",
			funcName:  "CreateInvoice",
			line:      10,
			imports:   []string{"fmt", "strings", "os", "io"},
			wantCount: 0,
		},
		{
			name:      "database_sql_forbidden",
			pkg:       "billing",
			funcName:  "CreateInvoice",
			line:      5,
			imports:   []string{"database/sql"},
			wantCount: 1,
		},
		{
			name:      "net_http_forbidden",
			pkg:       "user",
			funcName:  "Login",
			line:      8,
			imports:   []string{"net/http"},
			wantCount: 1,
		},
		{
			name:      "grpc_forbidden",
			pkg:       "order",
			funcName:  "Process",
			line:      3,
			imports:   []string{"google.golang.org/grpc"},
			wantCount: 1,
		},
		{
			name:      "subpackage_of_forbidden_also_caught",
			pkg:       "order",
			funcName:  "Process",
			line:      3,
			imports:   []string{"database/sql/driver", "net/http/httptest"},
			wantCount: 2,
		},
		{
			name:      "mixed_allowed_and_forbidden",
			pkg:       "billing",
			funcName:  "CreateInvoice",
			line:      10,
			imports:   []string{"fmt", "database/sql", "strings", "net/http"},
			wantCount: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runCollectForbiddenImports(t, c)
		})
	}
}
