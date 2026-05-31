//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgHelpersZeroCov — classifyFieldArgLocation / collectFieldArgSlices / opModelName / collectQueryMethods 직접 커버

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestClassifyFieldArgLocation_ZeroCov(t *testing.T) {
	pathParams := map[string]bool{"id": true}
	queryParams := map[string]bool{"cursor": true}

	cases := []struct {
		fa   FieldArg
		want ParamLocation
	}{
		{FieldArg{Literal: "5"}, LocLiteral},
		{FieldArg{Source: "currentUser", Field: "ID"}, LocUser},
		{FieldArg{Source: "request", Field: "id"}, LocPath},
		{FieldArg{Source: "request", Field: "cursor"}, LocQuery},
		{FieldArg{Source: "request", Field: "name"}, LocBody},
		{FieldArg{Source: "course", Field: "ID"}, LocVar},
		{FieldArg{}, ""}, // no source, no literal → untouched
	}
	for i, c := range cases {
		fa := c.fa
		classifyFieldArgLocation(&fa, pathParams, queryParams)
		if fa.Location != c.want {
			t.Errorf("case %d: Location = %q, want %q", i, fa.Location, c.want)
		}
	}
}

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

func TestCollectQueryMethods_ZeroCov(t *testing.T) {
	seqs := []ssac.Sequence{
		{Type: ssac.SeqGet, Model: "Course.FindByID", Package: "queries"},
		{Type: ssac.SeqGet, Model: "Course.FindByID", Package: "queries"}, // dup → skipped
		{Type: ssac.SeqPost, Model: "Order.Insert", Package: "session"},
		{Type: ssac.SeqResponse, Model: "ignored"}, // non-CRUD → skipped
	}
	methods := collectQueryMethods(seqs)
	if len(methods) != 2 {
		t.Fatalf("expected 2 query methods, got %d: %+v", len(methods), methods)
	}
	if methods[0].Name != "CourseFindByID" {
		t.Errorf("method[0].Name = %q, want CourseFindByID", methods[0].Name)
	}
	if methods[1].Package != "session" {
		t.Errorf("method[1].Package = %q, want session", methods[1].Package)
	}
}
