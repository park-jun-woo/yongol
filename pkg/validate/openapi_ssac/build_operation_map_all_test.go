//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what TestBuildOperationMapAll — buildOperationMapAll/MethodMapAll 이 단일 사이트와 멀티 도메인 모두에서 전체 operationId 를 합치는지 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildOperationMapAll(t *testing.T) {
	t.Run("single-site sees the singular doc", func(t *testing.T) {
		fs := &yongol.Fullstack{OpenAPIDoc: opDoc("/users", "CreateUser")}
		if _, ok := buildOperationMapAll(fs)["CreateUser"]; !ok {
			t.Fatal("single-site op missing")
		}
		if _, ok := buildOperationMethodMapAll(fs)["CreateUser"]; !ok {
			t.Fatal("single-site method-map op missing")
		}
	})

	t.Run("domain mode merges ops across all domains", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &manifest.ProjectConfig{Domains: map[string]manifest.DomainConfig{
				"public": {}, "admin": {},
			}},
			DomainOpenAPIDocs: map[string]*openapi3.T{
				"public": opDoc("/users", "CreateUser"),
				"admin":  opDoc("/admin/users", "CreateAdminUser"),
			},
		}
		m := buildOperationMapAll(fs)
		if _, ok := m["CreateUser"]; !ok {
			t.Error("public domain op missing from buildOperationMapAll")
		}
		if _, ok := m["CreateAdminUser"]; !ok {
			t.Error("admin domain op missing from buildOperationMapAll")
		}
		mm := buildOperationMethodMapAll(fs)
		if _, ok := mm["CreateUser"]; !ok {
			t.Error("public domain op missing from buildOperationMethodMapAll")
		}
		if _, ok := mm["CreateAdminUser"]; !ok {
			t.Error("admin domain op missing from buildOperationMethodMapAll")
		}
	})
}
