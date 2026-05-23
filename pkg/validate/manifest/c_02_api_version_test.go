//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c02APIVersion — manifest apiVersion가 "yongol/v1"인지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC02APIVersion(t *testing.T) {
	cases := []TestC02APIVersionCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name:      "correct_version",
			fs:        &yongol.Fullstack{Manifest: &pm.ProjectConfig{APIVersion: "yongol/v1"}},
			wantCount: 0,
		},
		{
			name:      "wrong_version",
			fs:        &yongol.Fullstack{Manifest: &pm.ProjectConfig{APIVersion: "yongol/v2"}},
			wantCount: 1,
		},
		{
			name:      "empty_version",
			fs:        &yongol.Fullstack{Manifest: &pm.ProjectConfig{APIVersion: ""}},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC02APIVersion(t, c)
		})
	}
}
