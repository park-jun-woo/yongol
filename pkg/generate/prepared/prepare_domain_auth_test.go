//ff:func feature=generate type=test control=sequence
//ff:what DomainAuthFor — 도메인별 auth_mode override/상속 + 단일 사이트 nil 검증

package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDomainAuthFor(t *testing.T) {
	if got := DomainAuthFor(nil); got != nil {
		t.Fatalf("nil fs → want nil, got %v", got)
	}
	single := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if got := DomainAuthFor(single); got != nil {
		t.Fatalf("single-site → want nil, got %v", got)
	}

	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{Auth: &manifest.Auth{Mode: "cookie", Type: "session"}},
		Domains: map[string]manifest.DomainConfig{
			"public": {RoutePrefix: "/api"},                           // inherits cookie
			"admin":  {RoutePrefix: "/api/admin", AuthMode: "bearer"}, // override bearer
		},
	}}
	got := DomainAuthFor(fs)
	if got["public"].Mode != "cookie" || !got["public"].CsrfRequired {
		t.Errorf("public: want cookie+csrf, got mode=%q csrf=%v", got["public"].Mode, got["public"].CsrfRequired)
	}
	if got["admin"].Mode != "bearer" || got["admin"].CsrfRequired {
		t.Errorf("admin: want bearer no-csrf, got mode=%q csrf=%v", got["admin"].Mode, got["admin"].CsrfRequired)
	}
}
