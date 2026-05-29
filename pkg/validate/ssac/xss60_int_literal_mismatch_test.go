//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestXss60_IntLiteral_Mismatch — 정수 리터럴 publish + string subscribe 시 warning 발생 확인

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60_IntLiteral_Mismatch(t *testing.T) {
	publisher := parsessac.ServiceFunc{
		Name:     "SetTimeout",
		FileName: "set_timeout.ssac",
		Sequences: []parsessac.Sequence{
			{
				Type:  "publish",
				Topic: "config.timeout",
				Inputs: map[string]string{
					"Timeout": "86400",
				},
			},
		},
	}
	subscriber := parsessac.ServiceFunc{
		Name:     "OnConfigTimeout",
		FileName: "on_config_timeout.ssac",
		Subscribe: &parsessac.SubscribeInfo{
			Topic:       "config.timeout",
			MessageType: "TimeoutMsg",
		},
		Param: &parsessac.ParamInfo{TypeName: "TimeoutMsg"},
		Structs: []parsessac.StructInfo{{
			Name: "TimeoutMsg",
			Fields: []parsessac.StructField{
				{Name: "Timeout", Type: "string"},
			},
		}},
	}
	fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, nil)
	diags := xss60SubscribeFieldTypes(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "int64") {
		t.Errorf("expected int64 in message: %s", diags[0].Message)
	}
}
