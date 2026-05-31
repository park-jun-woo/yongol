//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestFieldArgHelpersZeroCov — classifyFieldArgLocation / collectFieldArgSlices / opModelName / collectQueryMethods 직접 커버
package ir

import (
	"testing"
)

func TestOpModelName_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		op   *Op
		want string
	}{
		{"get", &Op{Kind: OpGet, Get: &GetOp{Model: "Course"}}, "Course"},
		{"post", &Op{Kind: OpPost, Post: &PostOp{Model: "Order"}}, "Order"},
		{"put", &Op{Kind: OpPut, Put: &PutOp{Model: "Item"}}, "Item"},
		{"delete", &Op{Kind: OpDelete, Delete: &DeleteOp{Model: "Tag"}}, "Tag"},
		{"get-nil", &Op{Kind: OpGet}, ""},
		{"auth-default", &Op{Kind: OpAuth, Auth: &AuthOp{}}, ""},
	}
	for _, c := range cases {
		if got := opModelName(c.op); got != c.want {
			t.Errorf("%s: opModelName = %q, want %q", c.name, got, c.want)
		}
	}
}
