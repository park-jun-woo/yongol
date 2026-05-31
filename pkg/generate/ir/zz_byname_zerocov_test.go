//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	pddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestEnrichFieldArgLocations_ZeroCov(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{Args: []FieldArg{
			{Key: "id", Source: "request", Field: ".ID"},
			{Source: "currentUser", Field: ".OrgID"},
		}}},
	}
	enrichFieldArgLocations(ops, map[string]bool{".ID": true}, map[string]bool{})
	if ops[0].Get.Args[0].Location != LocPath {
		t.Errorf("expected LocPath, got %q", ops[0].Get.Args[0].Location)
	}
	if ops[0].Get.Args[1].Location != LocUser {
		t.Errorf("expected LocUser, got %q", ops[0].Get.Args[1].Location)
	}
}

func TestEnrichFieldArgDDL_ZeroCov(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{Model: "User", Args: []FieldArg{
			{Key: "ID", Field: ".OrgID"},
		}}},
	}
	fs := &yongol.Fullstack{
		DDLTables: []pddl.Table{{
			Name:       "users",
			PrimaryKey: []string{"id"},
			Columns:    map[string]pddl.Column{"id": {Name: "id"}},
		}},
	}
	enrichFieldArgDDL(ops, fs)
	if ops[0].Get.Args[0].SourceColumn != "org_id" {
		t.Errorf("SourceColumn=%q want org_id", ops[0].Get.Args[0].SourceColumn)
	}
	// nil fs -> only Pass1 runs, no panic
	enrichFieldArgDDL(ops, nil)
}

func TestExtractBodyFields_ZeroCov(t *testing.T) {
	schema := openapi3.NewObjectSchema()
	schema.Properties = openapi3.Schemas{
		"name": openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
	}
	schema.Required = []string{"name"}
	op := &openapi3.Operation{
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(schema),
		},
	}
	fields := extractBodyFields(op)
	if len(fields) != 1 || fields[0].Name != "name" || !fields[0].Required {
		t.Errorf("body fields wrong: %#v", fields)
	}
	// no request body -> nil
	if got := extractBodyFields(&openapi3.Operation{}); got != nil {
		t.Errorf("no body should be nil")
	}
}

func TestExtractOpenAPIParams_ZeroCov(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{
					OperationID: "GetCourse",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "id", In: "path", Required: true,
							Schema: openapi3.NewSchemaRef("", openapi3.NewIntegerSchema()),
						}},
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "cursor", In: "query",
							Schema: openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
						}},
					},
				},
			}),
		),
	}
	plan := &ServicePlan{OperationID: "GetCourse"}
	pp, qp := extractOpenAPIParams(&yongol.Fullstack{OpenAPIDoc: doc}, "GetCourse", plan)
	if !pp["id"] || !qp["cursor"] {
		t.Errorf("params not classified: path=%v query=%v", pp, qp)
	}
	// nil fs -> empty maps, no panic
	pp2, _ := extractOpenAPIParams(nil, "x", plan)
	if len(pp2) != 0 {
		t.Errorf("nil fs should give empty path params")
	}
}

func TestApplySecurityHeadersOverrides_ZeroCov(t *testing.T) {
	cfg := &SecurityHeadersConfig{Profile: "production"}
	sh := &manifest.SecurityHeadersConfig{
		Profile:        "dev",
		HSTS:           &manifest.HSTSConfig{MaxAge: 100, IncludeSubDomains: true, Preload: true},
		CSP:            &manifest.CSPConfig{ReportOnly: true, Directives: map[string][]string{"default-src": {"'self'"}}},
		XFrameOptions:  "DENY",
		ReferrerPolicy: "no-referrer",
	}
	applySecurityHeadersOverrides(cfg, sh)
	if cfg.Profile != "dev" || cfg.HSTSMaxAge != 100 || cfg.XFrameOptions != "DENY" {
		t.Errorf("overrides not applied: %#v", cfg)
	}
	if !cfg.CSPReportOnly {
		t.Errorf("dev profile should force CSPReportOnly")
	}
}
