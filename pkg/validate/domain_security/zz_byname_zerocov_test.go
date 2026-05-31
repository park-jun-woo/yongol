//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package domain_security

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func emptySecurityOp(opID string) *openapi3.Operation {
	sr := openapi3.SecurityRequirements{}
	return &openapi3.Operation{OperationID: opID, Security: &sr}
}

func securedOp(opID string) *openapi3.Operation {
	sr := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearer": {}}}
	return &openapi3.Operation{OperationID: opID, Security: &sr}
}

func TestByNamePathSecurityChecks_ZeroCov(t *testing.T) {
	adminItem := &openapi3.PathItem{Get: emptySecurityOp("AdminList")}
	if d := checkAdminPathSecurity("/admin/x", adminItem, "admin.yaml"); len(d) == 0 {
		t.Errorf("checkAdminPathSecurity expected diagnostic for empty security")
	}

	internalItem := &openapi3.PathItem{Post: securedOp("InternalCreate")}
	if d := checkInternalPathSecurity("/internal/x", internalItem, "internal.yaml"); len(d) == 0 {
		t.Errorf("checkInternalPathSecurity expected warning for secured internal op")
	}
}

func TestByNameOpDomainMap_ZeroCov(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "ListItems"}}),
	)}
	docs := []domainDoc{{Name: "core", Doc: doc, Cfg: manifest.DomainConfig{OpenAPI: "core.yaml"}}}

	m := buildOpDomainMap(docs)
	if m["ListItems"] != "core" {
		t.Errorf("buildOpDomainMap = %v", m)
	}

	result := map[string]string{}
	collectDocOpDomains(docs[0], result)
	if result["ListItems"] != "core" {
		t.Errorf("collectDocOpDomains = %v", result)
	}

	// checkUnconsumedOps: ListItems not consumed → diagnostic.
	consumed := map[string]struct{}{}
	if d := checkUnconsumedOps(docs[0], consumed, "XMO-21", "core"); len(d) == 0 {
		t.Errorf("checkUnconsumedOps expected diagnostic for unconsumed op")
	}
	// consumed → no diagnostic.
	consumed["ListItems"] = struct{}{}
	if d := checkUnconsumedOps(docs[0], consumed, "XMO-21", "core"); len(d) != 0 {
		t.Errorf("checkUnconsumedOps unexpected diagnostics: %v", d)
	}
}

func TestByNameCrossDomainChecks_ZeroCov(t *testing.T) {
	opDomain := map[string]string{"OtherOp": "other"}

	fetch := stml.FetchBlock{
		OperationID:   "OtherOp",
		NestedFetches: []stml.FetchBlock{{OperationID: "OtherOp"}},
	}
	if d := checkFetchCrossDomain(fetch, "core/page.html", "core", opDomain); len(d) == 0 {
		t.Errorf("checkFetchCrossDomain expected cross-domain warning")
	}

	page := stml.PageSpec{
		FileName: "core/page.html",
		Fetches:  []stml.FetchBlock{fetch},
		Actions:  []stml.ActionBlock{{OperationID: "OtherOp"}},
	}
	if d := checkPageCrossDomain(page, "core", opDomain); len(d) == 0 {
		t.Errorf("checkPageCrossDomain expected cross-domain warnings")
	}
}

func TestByNameConsumedOps_ZeroCov(t *testing.T) {
	pages := []stml.PageSpec{{
		Fetches: []stml.FetchBlock{{
			OperationID:   "ListItems",
			NestedFetches: []stml.FetchBlock{{OperationID: "ListSub"}},
		}},
		Actions: []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}
	consumed := collectConsumedOpsFromPages(pages)
	for _, op := range []string{"ListItems", "ListSub", "CreateItem"} {
		if _, ok := consumed[op]; !ok {
			t.Errorf("collectConsumedOpsFromPages missing %q", op)
		}
	}

	out := map[string]struct{}{}
	collectFetchOpsRecursive(pages[0].Fetches[0], out)
	if _, ok := out["ListSub"]; !ok {
		t.Errorf("collectFetchOpsRecursive missing nested op")
	}
}

func TestByNameRegoDeleteActions_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{{
			Rules: []rego.AllowRule{
				{Resource: "items", Actions: []string{"delete"}},
				{Resource: "users", Actions: []string{"read"}},
			},
		}},
	}
	result := collectRegoDeleteActions(fs)
	if _, ok := result["items"]; !ok {
		t.Errorf("collectRegoDeleteActions missing delete resource")
	}
}

// TestByNameLoadAndXMO_ZeroCov drives loadDomainOpenAPIDocs / xmo21 / xmo22 by
// name using real OpenAPI files on disk and a two-domain manifest.
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
