//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBodyLimitConfig_ZeroCov(t *testing.T) {
	// nil fullstack → defaults.
	if c := buildBodyLimitConfig(nil); c.BodyLimit != 1048576 {
		t.Errorf("nil default body limit = %d", c.BodyLimit)
	}
	// nil manifest.HTTP → defaults.
	fsNoHTTP := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if c := buildBodyLimitConfig(fsNoHTTP); c.MultipartLimit != 33554432 {
		t.Errorf("default multipart = %d", c.MultipartLimit)
	}
	// HTTP with explicit limits → parsed overrides.
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				HTTP: &manifest.HTTPConfig{
					BodyLimit:      "2MiB",
					MultipartLimit: "64MiB",
				},
			},
		},
	}
	c := buildBodyLimitConfig(fs)
	if c.BodyLimit != 2*1048576 {
		t.Errorf("body limit = %d, want %d", c.BodyLimit, 2*1048576)
	}
	if c.MultipartLimit != 64*1048576 {
		t.Errorf("multipart limit = %d, want %d", c.MultipartLimit, 64*1048576)
	}
}
