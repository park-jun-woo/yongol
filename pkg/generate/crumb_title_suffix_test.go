//ff:func feature=generate type=test control=sequence
//ff:what TestCrumbTitleSuffix — 앱명 결합 꼬리 산출 / manifest·앱명 부재 빈 문자열 검증

package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCrumbTitleSuffix(t *testing.T) {
	t.Run("app name joins with the static-title separator", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
		fs.Manifest.Metadata.Name = "zenflow"
		if got := crumbTitleSuffix(fs); got != " · zenflow" {
			t.Errorf("crumbTitleSuffix = %q, want \" · zenflow\"", got)
		}
	})

	t.Run("nil fs, nil manifest and empty name yield empty", func(t *testing.T) {
		if got := crumbTitleSuffix(nil); got != "" {
			t.Errorf("crumbTitleSuffix(nil) = %q, want empty", got)
		}
		if got := crumbTitleSuffix(&yongol.Fullstack{}); got != "" {
			t.Errorf("crumbTitleSuffix(no manifest) = %q, want empty", got)
		}
		if got := crumbTitleSuffix(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); got != "" {
			t.Errorf("crumbTitleSuffix(no app name) = %q, want empty", got)
		}
	})
}
