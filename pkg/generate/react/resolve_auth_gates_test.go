//ff:func feature=gen-react type=test control=sequence
//ff:what resolveAuthGates — manifest 부재/bearer/cookie 모드별 인증 게이트 해석 검증
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveAuthGates(t *testing.T) {
	// nil fullstack -> no auth gates
	if has, bearer, store := resolveAuthGates(nil); has || bearer || store != "" {
		t.Errorf("nil fs = (%v,%v,%q), want (false,false,\"\")", has, bearer, store)
	}

	// manifest present but backend.auth absent -> no auth gates
	noAuth := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if has, bearer, store := resolveAuthGates(noAuth); has || bearer || store != "" {
		t.Errorf("no backend.auth = (%v,%v,%q), want (false,false,\"\")", has, bearer, store)
	}

	// bearer mode + explicit memory store -> hasAuth, bearerAuth, memory
	bearerFs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend:  manifest.Backend{Auth: &manifest.Auth{Mode: "bearer"}},
		Frontend: manifest.Frontend{Auth: &manifest.FrontendAuth{Store: "memory"}},
	}}
	if has, bearer, store := resolveAuthGates(bearerFs); !has || !bearer || store != "memory" {
		t.Errorf("bearer = (%v,%v,%q), want (true,true,\"memory\")", has, bearer, store)
	}

	// cookie mode (defaulted) + frontend.auth absent -> hasAuth, not bearer,
	// store defaults to localStorage
	cookieFs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{Auth: &manifest.Auth{}},
	}}
	if has, bearer, store := resolveAuthGates(cookieFs); !has || bearer || store != "localStorage" {
		t.Errorf("cookie = (%v,%v,%q), want (true,false,\"localStorage\")", has, bearer, store)
	}
}
