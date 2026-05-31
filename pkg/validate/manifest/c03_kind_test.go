//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c03Kind — manifest kind가 "Project"인지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC03Kind(t *testing.T) {
	cases := []TestC03KindCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "correct_kind", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Kind: "Project"}}, wantCount: 0},
		{name: "wrong_kind", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Kind: "Service"}}, wantCount: 1},
		{name: "empty_kind", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Kind: ""}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC03Kind(t, c)
		})
	}
}
