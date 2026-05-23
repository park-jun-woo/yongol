//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what opLabel — operation의 reader-friendly 식별자 선택 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpLabel(t *testing.T) {
	cases := []struct {
		name string
		op   *openapi3.Operation
		want string
	}{
		{name: "nil_op", op: nil, want: "op"},
		{name: "with_operation_id", op: &openapi3.Operation{OperationID: "getUser"}, want: "getUser"},
		{name: "no_id_with_tag", op: &openapi3.Operation{Tags: []string{"Users"}}, want: "Users"},
		{name: "no_id_no_tag", op: &openapi3.Operation{}, want: "op"},
		{name: "id_takes_precedence_over_tag", op: &openapi3.Operation{OperationID: "getUser", Tags: []string{"Users"}}, want: "getUser"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := opLabel(c.op)
			if got != c.want {
				t.Errorf("opLabel(...) = %q, want %q", got, c.want)
			}
		})
	}
}
