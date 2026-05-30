package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환

func TestNestifyPath_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"/users/{id}":            "/users/:id",
		"/a/{x}/b/{y}":           "/a/:x/b/:y",
		"/static":                "/static",
		"/broken/{unterminated": "/broken/{unterminated",
	}
	for in, want := range cases {
		if got := nestifyPath(in); got != want {
			t.Errorf("nestifyPath(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestLcFirst_ZeroCov — 첫 글자 소문자

func TestNestLcFirst_ZeroCov(t *testing.T) {
	if nestLcFirst("") != "" {
		t.Error("empty case")
	}
	if nestLcFirst("Charge") != "charge" {
		t.Error("Charge case")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestContainsOpKind_ZeroCov — OpKind 포함 여부

func TestContainsOpKind_ZeroCov(t *testing.T) {
	ops := []ir.Op{{Kind: ir.OpAuth}, {Kind: ir.OpGet}}
	if !containsOpKind(ops, ir.OpAuth) {
		t.Error("expected OpAuth present")
	}
	if containsOpKind(ops, ir.OpPublish) {
		t.Error("OpPublish should be absent")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestEnsurePkgMap_ZeroCov — 서브맵 초기화

func TestEnsurePkgMap_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	ensurePkgMap(pm, "billing")
	if pm["billing"] == nil {
		t.Error("expected submap created")
	}
	// idempotent
	pm["billing"]["x"] = true
	ensurePkgMap(pm, "billing")
	if !pm["billing"]["x"] {
		t.Error("ensurePkgMap should not reset existing submap")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAddOpPackageRef_ZeroCov — @call/@eval 패키지 참조 추가

func TestAddOpPackageRef_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}})
	addOpPackageRef(pm, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "rules", Function: "IsExpired"}})
	// no-package call → ignored
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: ""}})
	// other kind → ignored
	addOpPackageRef(pm, ir.Op{Kind: ir.OpGet})
	if !pm["billing"]["Charge"] {
		t.Error("expected billing.Charge")
	}
	if !pm["rules"]["IsExpired"] {
		t.Error("expected rules.IsExpired")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectOpsPackages_ZeroCov — Op 배열 패키지 수집

func TestCollectOpsPackages_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	ops := []ir.Op{
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
		{Kind: ir.OpGet},
	}
	collectOpsPackages(pm, ops)
	if !pm["billing"]["Charge"] {
		t.Error("expected billing.Charge collected")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildSortedPackages_ZeroCov — 정렬된 externalPackage 슬라이스

func TestBuildSortedPackages_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{
		"zeta":  {"B": true, "A": true},
		"alpha": {"X": true},
	}
	got := buildSortedPackages(pm)
	if len(got) != 2 || got[0].Name != "alpha" || got[1].Name != "zeta" {
		t.Fatalf("unexpected order: %+v", got)
	}
	if len(got[1].Methods) != 2 || got[1].Methods[0] != "A" || got[1].Methods[1] != "B" {
		t.Errorf("methods not sorted: %+v", got[1].Methods)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasPublishPlans_ZeroCov — @publish Op 존재 여부

func TestHasPublishPlans_ZeroCov(t *testing.T) {
	with := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpPublish}}}}}
	if !hasPublishPlans(with) {
		t.Error("expected publish present")
	}
	without := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpGet}}}}}
	if hasPublishPlans(without) {
		t.Error("publish should be absent")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasAuthPlans_ZeroCov — @auth Op 존재 여부

func TestHasAuthPlans_ZeroCov(t *testing.T) {
	with := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpAuth}}}}}
	if !hasAuthPlans(with) {
		t.Error("expected auth present")
	}
	without := map[string][]*ir.ServicePlan{"f": {{Ops: []ir.Op{{Kind: ir.OpGet}}}}}
	if hasAuthPlans(without) {
		t.Error("auth should be absent")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestResolveProjectID_ZeroCov — manifest name → projectID, fallback "app"

func TestResolveProjectID_ZeroCov(t *testing.T) {
	if got := resolveProjectID(&yongol.Fullstack{}); got != "app" {
		t.Errorf("fallback = %q", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "myapp"}}}
	if got := resolveProjectID(fs); got != "myapp" {
		t.Errorf("resolveProjectID = %q", got)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteScaffold_ZeroCov — scaffold 3종 파일 기록

func TestWriteScaffold_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeScaffold(dir, "myapp"); err != nil {
		t.Fatalf("writeScaffold error: %v", err)
	}
	for _, name := range []string{"package.json", "tsconfig.json", "nest-cli.json"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteAuthzModule_ZeroCov — authz 모듈/서비스 파일 기록

func TestWriteAuthzModule_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeAuthzModule(dir); err != nil {
		t.Fatalf("writeAuthzModule error: %v", err)
	}
	for _, name := range []string{"authz/authz.service.ts", "authz/authz.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteQueueModule_ZeroCov — queue 모듈/서비스 파일 기록

func TestWriteQueueModule_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeQueueModule(dir); err != nil {
		t.Fatalf("writeQueueModule error: %v", err)
	}
	for _, name := range []string{"queue/queue.service.ts", "queue/queue.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWritePrismaModule_ZeroCov — prisma 모듈/서비스 파일 기록

func TestWritePrismaModule_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writePrismaModule(dir); err != nil {
		t.Fatalf("writePrismaModule error: %v", err)
	}
	for _, name := range []string{"prisma/prisma.service.ts", "prisma/prisma.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteFuncStubs_ZeroCov — 외부 패키지 stub 파일 기록

func TestWriteFuncStubs_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	pkgs := []externalPackage{{Name: "billing", Methods: []string{"Charge"}}}
	if err := writeFuncStubs(dir, pkgs); err != nil {
		t.Fatalf("writeFuncStubs error: %v", err)
	}
	for _, name := range []string{"billing/billing.service.ts", "billing/billing.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectExternalPackages_ZeroCov — ServicePlan 맵 → 외부 패키지 목록

func TestCollectExternalPackages_ZeroCov(t *testing.T) {
	plansByFeature := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
			{Kind: ir.OpGet},
		}}},
	}
	got := collectExternalPackages(plansByFeature)
	if len(got) != 1 || got[0].Name != "billing" || got[0].Methods[0] != "Charge" {
		t.Errorf("unexpected: %+v", got)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestResolveHTTPRoute_ZeroCov — operationId → HTTP method/path 해석 + nil 분기

func TestResolveHTTPRoute_ZeroCov(t *testing.T) {
	// nil doc → no-op
	plan := &ir.ServicePlan{OperationID: "getUser"}
	resolveHTTPRoute(plan, &yongol.Fullstack{})
	if plan.HTTPMethod != "" {
		t.Error("expected no route with nil doc")
	}
	// matching doc → resolves method + path
	op := &openapi3.Operation{OperationID: "getUser"}
	paths := openapi3.NewPaths(openapi3.WithPath("/users/{id}", &openapi3.PathItem{Get: op}))
	fs := &yongol.Fullstack{OpenAPIDoc: &openapi3.T{Paths: paths}}
	resolveHTTPRoute(plan, fs)
	if plan.HTTPMethod != "GET" || plan.URLPath != "/users/:id" {
		t.Errorf("resolveHTTPRoute = %q %q", plan.HTTPMethod, plan.URLPath)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWritePrismaSchema_ZeroCov — empty(skip) + DDL → schema.prisma 기록

func TestWritePrismaSchema_ZeroCov(t *testing.T) {
	// empty → no-op
	if err := writePrismaSchema(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Fatalf("empty writePrismaSchema error: %v", err)
	}
	// with tables → schema file written
	backend := t.TempDir()
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{{
		Name:        "users",
		ColumnOrder: []string{"id"},
		Columns:     map[string]ddl.Column{"id": {RawType: "BIGINT", NotNull: true}},
		PrimaryKey:  []string{"id"},
	}}}
	if err := writePrismaSchema(fs, backend); err != nil {
		t.Fatalf("writePrismaSchema error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(backend, "prisma", "schema.prisma")); err != nil {
		t.Errorf("expected schema.prisma: %v", err)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteBootFiles_ZeroCov — main.ts + app.module.ts 파일 기록

func TestWriteBootFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	plan := &ir.BootPlan{ProjectID: "myapp"}
	if err := writeBootFiles(dir, plan, []string{"users"}, []string{"queue"}); err != nil {
		t.Fatalf("writeBootFiles error: %v", err)
	}
	for _, name := range []string{"main.ts", "app.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
