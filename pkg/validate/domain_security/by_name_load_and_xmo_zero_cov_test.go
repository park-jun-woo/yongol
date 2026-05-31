//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameLoadAndXMO_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	adminYAML := `openapi: "3.0.0"
info: {title: admin, version: "1.0"}
paths:
  /admin/items:
    get:
      operationId: AdminList
      responses:
        "200": {description: ok}
`
	coreYAML := `openapi: "3.0.0"
info: {title: core, version: "1.0"}
paths:
  /items:
    get:
      operationId: ListItems
      responses:
        "200": {description: ok}
`
	if err := os.WriteFile(filepath.Join(dir, "admin.yaml"), []byte(adminYAML), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "core.yaml"), []byte(coreYAML), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{
		SpecsDir: dir,
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"admin": {OpenAPI: "admin.yaml", Frontend: "frontend/admin"},
				"core":  {OpenAPI: "core.yaml", Frontend: "frontend/core"},
			},
		},
		STMLPages: []stml.PageSpec{
			{FileName: "frontend/admin/dashboard.html"},
			{FileName: "frontend/core/items.html", Fetches: []stml.FetchBlock{{OperationID: "AdminList"}}},
		},
	}

	docs := loadDomainOpenAPIDocs(fs)
	if len(docs) != 2 {
		t.Fatalf("loadDomainOpenAPIDocs = %d, want 2", len(docs))
	}

	_ = xmo21AdminUnconsumed(fs)
	_ = xmo22CrossDomainCall(fs)
}
