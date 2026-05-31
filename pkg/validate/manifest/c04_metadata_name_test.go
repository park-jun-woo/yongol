//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c04MetadataName — manifest metadata.name 비어있는지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC04MetadataName(t *testing.T) {
	cases := []TestC04MetadataNameCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "valid_name", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Metadata: pm.Metadata{Name: "myproject"}}}, wantCount: 0},
		{name: "empty_name", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Metadata: pm.Metadata{Name: ""}}}, wantCount: 1},
		{name: "whitespace_name", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Metadata: pm.Metadata{Name: "  "}}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC04MetadataName(t, c)
		})
	}
}
