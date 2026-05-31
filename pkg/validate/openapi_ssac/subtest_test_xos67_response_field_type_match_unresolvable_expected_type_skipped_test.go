//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchUnresolvableExpectedTypeSkipped — unresolvable expected type skipped 서브테스트
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchUnresolvableExpectedTypeSkipped(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name: "getUser",
				Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"name": `"hello"`}},
				},
			},
		},
	}
	g := &rule.Ground{Types: map[string]string{}}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d", len(diags))
	}

}
