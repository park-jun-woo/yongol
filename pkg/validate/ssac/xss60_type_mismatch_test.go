//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestXss60_TypeMismatch_Warning — publish/subscribe 타입 불일치 시 WARNING 발생 확인

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60_TypeMismatch_Warning(t *testing.T) {
	// @publish sends user.ID (BIGINT → int64), subscriber expects string. Warning.
	tables := []ddl.Table{{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id": {Name: "id", RawType: "BIGINT"},
		},
	}}
	publisher := parsessac.ServiceFunc{
		Name:     "CreateUser",
		FileName: "create_user.ssac",
		Sequences: []parsessac.Sequence{
			{
				Type:   "post",
				Result: &parsessac.Result{Type: "User", Var: "user"},
			},
			{
				Type:  "publish",
				Topic: "user.created",
				Inputs: map[string]string{
					"UserID": "user.ID",
				},
			},
		},
	}
	subscriber := parsessac.ServiceFunc{
		Name:     "OnUserCreated",
		FileName: "on_user_created.ssac",
		Subscribe: &parsessac.SubscribeInfo{
			Topic:       "user.created",
			MessageType: "UserCreatedMsg",
		},
		Param: &parsessac.ParamInfo{TypeName: "UserCreatedMsg"},
		Structs: []parsessac.StructInfo{{
			Name: "UserCreatedMsg",
			Fields: []parsessac.StructField{
				{Name: "UserID", Type: "string"},
			},
		}},
	}
	fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, tables)
	diags := xss60SubscribeFieldTypes(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %v", len(diags), diags)
	}
	d := diags[0]
	if d.Level != diagnostic.LevelWarning {
		t.Errorf("expected WARNING, got %v", d.Level)
	}
	if !strings.Contains(d.Message, "XSS-60") {
		t.Errorf("message should contain XSS-60: %s", d.Message)
	}
	if !strings.Contains(d.Message, "int64") || !strings.Contains(d.Message, "string") {
		t.Errorf("message should mention both types: %s", d.Message)
	}
}
