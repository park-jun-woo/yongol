//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestXss60_TypeMatch_NoDiag — publish/subscribe 타입 일치 시 진단 미발생 확인

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60_TypeMatch_NoDiag(t *testing.T) {
	// @publish sends user.ID (BIGINT → int64), subscriber expects int64. No warning.
	tables := []ddl.Table{{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":    {Name: "id", RawType: "BIGINT"},
			"email": {Name: "email", RawType: "TEXT"},
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
					"Email":  "user.Email",
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
				{Name: "UserID", Type: "int64"},
				{Name: "Email", Type: "string"},
			},
		}},
	}
	fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, tables)
	diags := xss60SubscribeFieldTypes(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics, got %d: %v", len(diags), diags)
	}
}
