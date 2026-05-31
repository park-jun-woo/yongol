//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what resolveSecurityHeaders — manifest.security_headers + production 기본값 병합
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveSecurityHeaders_Defaults(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{nil, {}, {Manifest: &pmanifest.ProjectConfig{}}} {
		out := resolveSecurityHeaders(fs)
		if out.Profile != defaultSHProfile {
			t.Errorf("profile = %q, want %q", out.Profile, defaultSHProfile)
		}
		if out.HSTSMaxAge != defaultHSTSMaxAge || !out.CSPEnabled {
			t.Errorf("default preset wrong: %+v", out)
		}
		if out.XFrameOptions != defaultXFrameOptions {
			t.Errorf("x-frame = %q, want %q", out.XFrameOptions, defaultXFrameOptions)
		}
	}
}
