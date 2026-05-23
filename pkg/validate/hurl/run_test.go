//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-structural
//ff:what Run — Hurl 검증 전체 실행 (H-1 + H-2) 통합 검증

package hurl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	cases := []TestRunCase{
		{
			name:      "clean_no_diags",
			setup:     func(dir string) {},
			presences: nil,
			wantCodes: nil,
		},
		{
			name: "deprecated_feature_triggers_h1",
			setup: func(dir string) {
				os.MkdirAll(filepath.Join(dir, "tests"), 0o755)
				os.WriteFile(filepath.Join(dir, "tests", "login.feature"), []byte(""), 0o644)
			},
			presences: nil,
			wantCodes: []string{"[H-1]"},
		},
		{
			name:  "declared_empty_triggers_h2",
			setup: func(dir string) {},
			presences: map[yongol.SSOTKind]yongol.SSOTPresence{
				yongol.KindScenario: yongol.SSOTDeclared,
			},
			wantCodes: []string{"[H-2]"},
		},
		{
			name: "both_rules_fire",
			setup: func(dir string) {
				os.MkdirAll(filepath.Join(dir, "tests"), 0o755)
				os.WriteFile(filepath.Join(dir, "tests", "login.feature"), []byte(""), 0o644)
			},
			presences: map[yongol.SSOTKind]yongol.SSOTPresence{
				yongol.KindScenario: yongol.SSOTDeclared,
			},
			wantCodes: []string{"[H-1]", "[H-2]"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runRun(t, c)
		})
	}
}
