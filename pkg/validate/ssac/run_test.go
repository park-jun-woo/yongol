//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what Run — SSaC 전체 검증 smoke test (빈 Fullstack → 0 diags, 단건 S-1 위반 검출)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("EmptyFullstack", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Errorf("Run(empty) produced %d diags, want 0; first: %s", len(diags), diags[0].Message)
		}
	})
	t.Run("SingleS1Violation", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					FileName: "test.ssac",
					Sequences: []ssac.Sequence{
						{Type: "get", Line: 5, Model: ""},
					},
				},
			},
		}
		diags := Run(fs)

		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-1]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Run() did not produce S-1 diagnostic for @get with empty Model; got %d diags", len(diags))
		}
	})
}
