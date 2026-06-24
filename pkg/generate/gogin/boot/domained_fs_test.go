//ff:func feature=gen-gogin type=test control=sequence
//ff:what domainedFS — 2 도메인(public/admin) Fullstack 테스트 픽스처 생성

package boot

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func domainedFS(middleware []string) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "example.com/app", Middleware: middleware},
			Domains: map[string]manifest.DomainConfig{
				"public": {RoutePrefix: "/api"},
				"admin":  {RoutePrefix: "/api/admin"},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{
			"public": {Paths: openapi3.NewPaths(
				openapi3.WithPath("/login", &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "Login"}}),
			)},
			"admin": {Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "AdminListUsers"}}),
			)},
		},
	}
}
