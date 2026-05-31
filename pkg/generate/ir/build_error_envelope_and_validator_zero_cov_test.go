//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildErrorEnvelopeAndValidator_ZeroCov(t *testing.T) {
	if c := buildErrorEnvelopeConfig(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); c == nil {
		t.Errorf("error envelope should always be non-nil")
	}
	if c := buildRequestValidatorConfig(); !c.Active {
		t.Errorf("request validator should be active")
	}
}
