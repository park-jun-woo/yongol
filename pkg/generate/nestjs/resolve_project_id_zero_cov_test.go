//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveProjectID_ZeroCov(t *testing.T) {
	if got := resolveProjectID(&yongol.Fullstack{}); got != "app" {
		t.Errorf("fallback = %q", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "myapp"}}}
	if got := resolveProjectID(fs); got != "myapp" {
		t.Errorf("resolveProjectID = %q", got)
	}
}
