//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildCORSConfig_ZeroCov(t *testing.T) {
	if c := buildCORSConfig(nil); c != nil {
		t.Errorf("nil → nil cors")
	}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				CORS: &manifest.CORSConfig{
					AllowOrigins:     []string{"https://x"},
					AllowCredentials: true,
				},
			},
		},
	}
	c := buildCORSConfig(fs)
	if c == nil || !c.AllowCredentials || len(c.AllowOrigins) != 1 {
		t.Errorf("cors config = %+v", c)
	}
}
