//ff:func feature=validate type=test control=sequence topic=ssac-statemachine
//ff:what TestXsm71StateInputTypeBranches — xsm71StateInputType 가드/skip 분기 검증
package ssac_statemachine

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXsm71StateInputTypeBranches(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		if d := xsm71StateInputType(fs); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("non-state sequences skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:      "Fn",
				Sequences: []parsessac.Sequence{{Type: "get"}, {Type: "post"}},
			}},
		}
		fs.SetGround(&rule.Ground{Types: map[string]string{}})
		if d := xsm71StateInputType(fs); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})
}
