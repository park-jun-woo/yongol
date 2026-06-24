//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestStepUsesPerDomainDoc — OpenAPI/STML 게이트 판정 분기 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStepUsesPerDomainDoc(t *testing.T) {
	cases := []struct {
		name  string
		kinds []yongol.SSOTKind
		want  bool
	}{
		{"openapi", []yongol.SSOTKind{yongol.KindOpenAPI}, true},
		{"stml", []yongol.SSOTKind{yongol.KindSTML, yongol.KindStates}, true},
		{"ddl_only", []yongol.SSOTKind{yongol.KindDDL}, false},
		{"config_only", []yongol.SSOTKind{yongol.KindConfig}, false},
		{"none", nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := step{Kinds: c.kinds}
			if got := s.usesPerDomainDoc(); got != c.want {
				t.Errorf("usesPerDomainDoc(%v) = %v, want %v", c.kinds, got, c.want)
			}
		})
	}
}
