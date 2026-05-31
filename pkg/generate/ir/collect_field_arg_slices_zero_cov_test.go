//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestFieldArgHelpersZeroCov — classifyFieldArgLocation / collectFieldArgSlices / opModelName / collectQueryMethods 직접 커버
package ir

import (
	"testing"
)

func TestCollectFieldArgSlices_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		op   *Op
		want int
	}{
		{"get", &Op{Kind: OpGet, Get: &GetOp{}}, 2},
		{"post", &Op{Kind: OpPost, Post: &PostOp{}}, 1},
		{"put", &Op{Kind: OpPut, Put: &PutOp{}}, 1},
		{"delete", &Op{Kind: OpDelete, Delete: &DeleteOp{}}, 1},
		{"auth", &Op{Kind: OpAuth, Auth: &AuthOp{}}, 1},
		{"state", &Op{Kind: OpState, State: &StateOp{}}, 1},
		{"call", &Op{Kind: OpCall, Call: &CallOp{}}, 1},
		{"eval", &Op{Kind: OpEval, Eval: &EvalOp{}}, 1},
		{"publish", &Op{Kind: OpPublish, Publish: &PublishOp{}}, 2},
		{"get-nil", &Op{Kind: OpGet}, 0},
		{"response-default", &Op{Kind: OpResponse}, 0},
	}
	for _, c := range cases {
		got := collectFieldArgSlices(c.op)
		if len(got) != c.want {
			t.Errorf("%s: got %d slices, want %d", c.name, len(got), c.want)
		}
	}
}
