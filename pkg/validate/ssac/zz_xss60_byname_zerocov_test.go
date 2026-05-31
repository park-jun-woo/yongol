//ff:func feature=validate type=test control=sequence
//ff:what TestXss60ByName_ZeroCov — XSS-60 publish/subscribe 타입 수집·비교 헬퍼 직접 호출

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss60CollectFieldTypes_ZeroCov(t *testing.T) {
	fields := map[string]string{}
	seq := parsessac.Sequence{
		Type:  "publish",
		Topic: "order.completed",
		Inputs: map[string]string{
			"status": `"done"`, // string literal
			"count":  "42",     // int literal
		},
		Fields: map[string]string{
			"note": `"x"`,
		},
	}
	fn := parsessac.ServiceFunc{Name: "Pub"}
	xss60CollectFieldTypes(fields, seq, fn, map[string]*ddl.Table{})
	if fields["status"] != "string" || fields["count"] != "int64" || fields["note"] != "string" {
		t.Errorf("collected field types = %v", fields)
	}
}

func TestXss60CollectPublishTypes_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
		Name: "Pub",
		Sequences: []parsessac.Sequence{
			{Type: "publish", Topic: "t.one", Inputs: map[string]string{"a": "1"}},
			{Type: "post", Topic: ""}, // skipped (not publish)
		},
	}}}
	out := xss60CollectPublishTypes(fs, map[string]*ddl.Table{})
	if out["t.one"]["a"] != "int64" {
		t.Errorf("publish types = %v", out)
	}
}

func TestXss60InferDotExpr_ZeroCov(t *testing.T) {
	fn := parsessac.ServiceFunc{
		Name: "Pub",
		Sequences: []parsessac.Sequence{
			{Type: "get", Result: &parsessac.Result{Var: "user", Type: "User"}},
		},
	}
	tables := map[string]*ddl.Table{
		"users": {Name: "users", Columns: map[string]ddl.Column{
			"email": {Name: "email", RawType: "TEXT"},
		}},
	}
	// var.Field resolved through model→table→column.
	got := xss60InferDotExpr("user.Email", 4, fn, tables)
	if got == "" {
		t.Errorf("expected Go type for user.Email, got empty")
	}
	// unresolvable variable model → "".
	if v := xss60InferDotExpr("ghost.X", 5, fn, tables); v != "" {
		t.Errorf("unknown var should yield empty, got %q", v)
	}
}

func TestXss60CheckSubscriber_ZeroCov(t *testing.T) {
	// no subscribe → nil.
	if d := xss60CheckSubscriber(parsessac.ServiceFunc{}, nil); d != nil {
		t.Errorf("non-subscriber should yield nil")
	}
	// subscriber with topic not in publish map → nil.
	fn := parsessac.ServiceFunc{
		Name:      "OnDone",
		Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
		Param:     &parsessac.ParamInfo{TypeName: "Msg"},
		Structs: []parsessac.StructInfo{{
			Name:   "Msg",
			Fields: []parsessac.StructField{{Name: "OrderID", Type: "int64"}},
		}},
	}
	if d := xss60CheckSubscriber(fn, map[string]map[string]string{}); d != nil {
		t.Errorf("topic absent from publishers should yield nil, got %v", d)
	}
	// matching topic, compatible field → no diag.
	pub := map[string]map[string]string{"order.completed": {"OrderID": "int64"}}
	if d := xss60CheckSubscriber(fn, pub); d != nil {
		t.Errorf("compatible field should yield no diag, got %v", d)
	}
}
