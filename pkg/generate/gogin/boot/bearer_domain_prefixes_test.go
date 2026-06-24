//ff:func feature=gen-gogin type=test control=sequence
//ff:what bearerDomainPrefixes — bearer 도메인 prefix 수집 + 단일 사이트 nil 검증

package boot

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBearerDomainPrefixes(t *testing.T) {
	if got := bearerDomainPrefixes(nil); got != nil {
		t.Fatalf("nil fs → want nil, got %v", got)
	}
	if got := bearerDomainPrefixes(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); got != nil {
		t.Fatalf("single-site → want nil, got %v", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{Auth: &manifest.Auth{Mode: "cookie"}},
		Domains: map[string]manifest.DomainConfig{
			"public": {RoutePrefix: "/api"},                           // cookie (inherits) → not exempt
			"admin":  {RoutePrefix: "/api/admin", AuthMode: "bearer"}, // bearer → exempt
		},
	}}
	if got := bearerDomainPrefixes(fs); !reflect.DeepEqual(got, []string{"/api/admin"}) {
		t.Errorf("want [/api/admin], got %v", got)
	}
}
