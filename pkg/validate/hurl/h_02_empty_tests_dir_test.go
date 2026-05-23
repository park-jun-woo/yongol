//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-structural
//ff:what h02EmptyTestsDir — SSOTDeclared/SSOTPopulated/SSOTAbsent 상태에 따른 H-2 진단 검증

package hurl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestH02EmptyTestsDir(t *testing.T) {
	cases := []TestH02EmptyTestsDirCase{
		{
			name: "declared_produces_warning",
			fs: &yongol.Fullstack{
				Presences: map[yongol.SSOTKind]yongol.SSOTPresence{
					yongol.KindScenario: yongol.SSOTDeclared,
				},
			},
			wantCount: 1,
		},
		{
			name:      "absent_no_diag",
			fs:        &yongol.Fullstack{},
			wantCount: 0,
		},
		{
			name:      "nil_fullstack_presences_no_diag",
			fs:        &yongol.Fullstack{Presences: nil},
			wantCount: 0,
		},
		{
			name: "populated_no_diag",
			fs: &yongol.Fullstack{
				Presences: map[yongol.SSOTKind]yongol.SSOTPresence{
					yongol.KindScenario: yongol.SSOTPopulated,
				},
			},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runH02EmptyTestsDir(t, c)
		})
	}
}
