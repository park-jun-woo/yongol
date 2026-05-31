//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPost_ZeroCov2(t *testing.T) {
	op := convertPost(ssac.Sequence{Model: "Course.Create", Inputs: map[string]string{"Title": "request.title"}, Result: &ssac.Result{Var: "c", Type: "Course", Wrapper: "Page"}})
	if op.Kind != OpPost || op.Post == nil {
		t.Fatalf("expected OpPost, got %+v", op)
	}
	if op.Post.Model != "Course" || op.Post.Method != "Create" {
		t.Errorf("model/method = %q/%q", op.Post.Model, op.Post.Method)
	}
	if !op.Post.IsList {
		t.Error("expected IsList for wrapper")
	}
	// nil result branch
	op = convertPost(ssac.Sequence{Model: "Course.Create"})
	if op.Post.VarName != "" {
		t.Errorf("expected empty VarName, got %q", op.Post.VarName)
	}
}
