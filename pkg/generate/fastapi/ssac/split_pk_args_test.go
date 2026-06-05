//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestSplitPKArgs — splitPKArgs IsPK 기준 where/data 인자 분리 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestSplitPKArgs(t *testing.T) {
	args := []ir.FieldArg{
		{Key: "id", IsPK: true},
		{Key: "title", IsPK: false},
		{Key: "org_id", IsPK: true},
	}
	where, data := splitPKArgs(args)
	if len(where) != 2 || where[0].Key != "id" || where[1].Key != "org_id" {
		t.Errorf("where = %v, want [id org_id]", where)
	}
	if len(data) != 1 || data[0].Key != "title" {
		t.Errorf("data = %v, want [title]", data)
	}

	w, d := splitPKArgs(nil)
	if w != nil || d != nil {
		t.Errorf("nil input: got where=%v data=%v, want nil,nil", w, d)
	}
}
