//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestXss60_LiteralInfer — 문자열/정수 리터럴 타입 추론 및 불일치 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60_LiteralInfer(t *testing.T) {
	t.Run("StringLiteral_Inferred", func(t *testing.T) {
		publisher := parsessac.ServiceFunc{
			Name:     "CompleteOrder",
			FileName: "complete_order.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:  "publish",
					Topic: "order.completed",
					Inputs: map[string]string{
						"Status": `"completed"`,
					},
				},
			},
		}
		subscriber := parsessac.ServiceFunc{
			Name:     "OnOrderCompleted",
			FileName: "on_order_completed.ssac",
			Subscribe: &parsessac.SubscribeInfo{
				Topic:       "order.completed",
				MessageType: "OrderCompletedMsg",
			},
			Param: &parsessac.ParamInfo{TypeName: "OrderCompletedMsg"},
			Structs: []parsessac.StructInfo{{
				Name: "OrderCompletedMsg",
				Fields: []parsessac.StructField{
					{Name: "Status", Type: "string"},
				},
			}},
		}
		fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, nil)
		diags := xss60SubscribeFieldTypes(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics for string literal match, got %d: %v", len(diags), diags)
		}
	})

	t.Run("IntLiteral_Inferred", func(t *testing.T) {
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
					{Name: "Timeout", Type: "int64"},
				},
			}},
		}
		fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, nil)
		diags := xss60SubscribeFieldTypes(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics for int literal match, got %d: %v", len(diags), diags)
		}
	})

}
