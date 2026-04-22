//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestBuildFuncAnnot — //ff:func 조립 검증 (table-driven)

package ffannot

import "testing"

func TestBuildFuncAnnot(t *testing.T) {
	cases := []struct {
		name string
		in   FuncAnnot
		want string
	}{
		{"sequence", FuncAnnot{Feature: "service", Type: "handler", Control: "sequence"},
			"//ff:func feature=service type=handler control=sequence"},
		{"iteration-default-dim", FuncAnnot{Feature: "service", Type: "handler", Control: "iteration"},
			"//ff:func feature=service type=handler control=iteration dimension=1"},
		{"iteration-dim-2", FuncAnnot{Feature: "service", Type: "handler", Control: "iteration", Dimension: 2},
			"//ff:func feature=service type=handler control=iteration dimension=2"},
		{"with-topic", FuncAnnot{Feature: "service", Type: "handler", Control: "sequence", Topic: "transaction-boundary"},
			"//ff:func feature=service type=handler control=sequence topic=transaction-boundary"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := BuildFuncAnnot(tc.in)
			if got != tc.want {
				t.Fatalf("BuildFuncAnnot() = %q, want %q", got, tc.want)
			}
		})
	}
}
