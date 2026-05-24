//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestXss60_EdgeCases — unresolvable/nil/empty/no-subscriber 경계 조건 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss60_EdgeCases(t *testing.T) {
	t.Run("Unresolvable_Skipped", func(t *testing.T) {
		publisher := parsessac.ServiceFunc{
			Name:     "DoStuff",
			FileName: "do_stuff.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:  "publish",
					Topic: "stuff.done",
					Inputs: map[string]string{
						"Payload": "someUnknownThing",
					},
				},
			},
		}
		subscriber := parsessac.ServiceFunc{
			Name:     "OnStuffDone",
			FileName: "on_stuff_done.ssac",
			Subscribe: &parsessac.SubscribeInfo{
				Topic:       "stuff.done",
				MessageType: "StuffDoneMsg",
			},
			Param: &parsessac.ParamInfo{TypeName: "StuffDoneMsg"},
			Structs: []parsessac.StructInfo{{
				Name: "StuffDoneMsg",
				Fields: []parsessac.StructField{
					{Name: "Payload", Type: "string"},
				},
			}},
		}
		fs := xss60FS([]parsessac.ServiceFunc{publisher, subscriber}, nil)
		diags := xss60SubscribeFieldTypes(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics for unresolvable expression, got %d: %v", len(diags), diags)
		}
	})

	t.Run("NilFullstack", func(t *testing.T) {
		diags := xss60SubscribeFieldTypes(nil)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics for nil Fullstack, got %d", len(diags))
		}
	})

	t.Run("EmptyFullstack", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xss60SubscribeFieldTypes(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics for empty Fullstack, got %d", len(diags))
		}
	})

	t.Run("NoSubscriber", func(t *testing.T) {
		publisher := parsessac.ServiceFunc{
			Name:     "CreateUser",
			FileName: "create_user.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:  "publish",
					Topic: "user.created",
					Inputs: map[string]string{
						"UserID": `"abc"`,
					},
				},
			},
		}
		fs := xss60FS([]parsessac.ServiceFunc{publisher}, nil)
		diags := xss60SubscribeFieldTypes(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diagnostics when no subscriber exists, got %d", len(diags))
		}
	})
}
