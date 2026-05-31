//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildActiveBlocks_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	ps := &prepared.State{}
	blocks := buildActiveBlocks(fs, ps)
	if len(blocks) != 25 {
		t.Errorf("expected 25 boot blocks, got %d", len(blocks))
	}
}
