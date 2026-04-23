//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestPrepareAuth — Mode 기본값 해석 케이스 (empty→cookie, jwt/hybrid 그대로)

package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPrepareAuth(t *testing.T) {
	cases := []struct {
		name     string
		auth     *pmanifest.Auth
		wantMode string
		wantOK   bool
	}{
		{name: "absent", auth: nil, wantOK: false},
		{name: "mode_empty_defaults_cookie", auth: &pmanifest.Auth{}, wantMode: "cookie", wantOK: true},
		{name: "mode_bearer", auth: &pmanifest.Auth{Mode: "bearer"}, wantMode: "bearer", wantOK: true},
		{name: "mode_hybrid", auth: &pmanifest.Auth{Mode: "hybrid"}, wantMode: "hybrid", wantOK: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: tc.auth}}}
			got := authFor(fs)
			if got.Present != tc.wantOK {
				t.Fatalf("Present=%v want %v", got.Present, tc.wantOK)
			}
			if tc.wantOK && got.Mode != tc.wantMode {
				t.Errorf("Mode=%q want %q", got.Mode, tc.wantMode)
			}
		})
	}
}
