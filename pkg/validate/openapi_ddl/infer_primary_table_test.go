//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what inferPrimaryTable — operationId/path 기반 테이블 추론 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestInferPrimaryTable(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Name: "users", Columns: map[string]ddl.Column{}},
			{Name: "orders", Columns: map[string]ddl.Column{}},
		},
	}

	tests := []struct {
		name string
		op   *openapi3.Operation
		path string
		want string
	}{
		{
			name: "from operationId getUser -> users",
			op:   &openapi3.Operation{OperationID: "getUser"},
			path: "/api/users/{id}",
			want: "users",
		},
		{
			name: "from operationId listOrders -> orders",
			op:   &openapi3.Operation{OperationID: "listOrders"},
			path: "/api/orders",
			want: "orders",
		},
		{
			name: "from path fallback",
			op:   &openapi3.Operation{OperationID: "doSomethingRandom"},
			path: "/api/users/{id}",
			want: "users",
		},
		{
			name: "nil op uses path",
			op:   nil,
			path: "/orders",
			want: "orders",
		},
		{
			name: "raw segment matches table directly",
			op:   &openapi3.Operation{OperationID: "doSomething"},
			path: "/api/users",
			want: "users",
		},
		{
			name: "no match returns empty",
			op:   &openapi3.Operation{OperationID: "doSomething"},
			path: "/api/health",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferPrimaryTable(fs, tt.op, tt.path)
			if got != tt.want {
				t.Errorf("inferPrimaryTable(%v, %q) = %q, want %q", tt.op, tt.path, got, tt.want)
			}
		})
	}
}
