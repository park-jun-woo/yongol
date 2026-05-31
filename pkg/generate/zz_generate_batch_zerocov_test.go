//ff:func feature=gen type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버

package generate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAppendFieldlessFromChild_ZeroCov(t *testing.T) {
	result := map[string]bool{}
	// action with no fields → recorded
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "Op1"}}, result)
	if !result["Op1"] {
		t.Error("expected Op1 recorded")
	}
	// action with fields → not recorded
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "Op2", Fields: []stmlparser.FieldBind{{Name: "x"}}}}, result)
	if result["Op2"] {
		t.Error("Op2 has fields, should not be recorded")
	}
	// nil action
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action"}, result)
	// fetch / state / static / each recurse
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "F"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "S"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "T"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "each", Each: &stmlparser.EachBlock{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "E"}}}}}, result)
	for _, k := range []string{"F", "S", "T", "E"} {
		if !result[k] {
			t.Errorf("expected %q recorded via recursion", k)
		}
	}
	// nil pointers in recursive kinds → no panic
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "fetch"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "state"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "static"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "each"}, result)
	// unknown kind
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "bind"}, result)
}

func TestApplyGenerateOptions_ZeroCov(t *testing.T) {
	cfg := &generateConfig{}
	called := false
	applyGenerateOptions(cfg, []GenerateOption{func(c *generateConfig) { called = true }})
	if !called {
		t.Error("expected hook called")
	}
	// empty hooks → no-op
	applyGenerateOptions(cfg, nil)
}

func TestBuildDefaultFieldConstraints_ZeroCov(t *testing.T) {
	// nil schema
	if buildDefaultFieldConstraints(nil) != nil {
		t.Error("nil schema should return nil")
	}
	// empty properties
	if buildDefaultFieldConstraints(&openapi3.Schema{}) != nil {
		t.Error("empty props should return nil")
	}
	// real schema with required + a nil-value prop skipped
	schema := &openapi3.Schema{
		Required: []string{"email"},
		Properties: openapi3.Schemas{
			"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			"bad":   &openapi3.SchemaRef{}, // nil Value → skipped
		},
	}
	got := buildDefaultFieldConstraints(schema)
	if got == nil || len(got) != 1 {
		t.Fatalf("expected 1 field, got %v", got)
	}
	if !got["email"].Required || got["email"].Type != "string" || got["email"].Format != "email" {
		t.Errorf("email constraint = %+v", got["email"])
	}
}

func TestBuildFallbackStringFields_ZeroCov(t *testing.T) {
	got := buildFallbackStringFields([]string{"a", "b"})
	if len(got) != 2 || got["a"].Type != "string" {
		t.Fatalf("buildFallbackStringFields = %v", got)
	}
}

func TestExtractRequestBodySchema_ZeroCov(t *testing.T) {
	// nil request body
	if extractRequestBodySchema(&openapi3.Operation{}) != nil {
		t.Error("nil body should return nil")
	}
	// application/json present
	jsonSchema := &openapi3.Schema{Type: &openapi3.Types{"object"}}
	op := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Value: jsonSchema}},
		},
	}}}
	if extractRequestBodySchema(op) != jsonSchema {
		t.Error("should return json schema")
	}
	// fallback content type
	xmlSchema := &openapi3.Schema{Type: &openapi3.Types{"string"}}
	op2 := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{
			"application/xml": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Value: xmlSchema}},
		},
	}}}
	if extractRequestBodySchema(op2) != xmlSchema {
		t.Error("should fall back to xml schema")
	}
	// content present but no usable schema
	op3 := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{}},
	}}}
	if extractRequestBodySchema(op3) != nil {
		t.Error("no schema → nil")
	}
}

func TestCopyFrontendComponents_ZeroCov(t *testing.T) {
	// empty specsDir → no-op
	if err := copyFrontendComponents("", t.TempDir()); err != nil {
		t.Fatalf("empty specsDir err: %v", err)
	}
	// missing frontend dir → no-op
	specs := t.TempDir()
	if err := copyFrontendComponents(specs, t.TempDir()); err != nil {
		t.Fatalf("missing frontend err: %v", err)
	}
	// frontend is a file not dir → no-op
	if err := os.WriteFile(filepath.Join(specs, "frontend"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyFrontendComponents(specs, t.TempDir()); err != nil {
		t.Fatalf("frontend-as-file err: %v", err)
	}
	// real frontend dir with a .tsx file → copied
	specs2 := t.TempDir()
	fe := filepath.Join(specs2, "frontend", "components")
	if err := os.MkdirAll(fe, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(fe, "Foo.tsx"), []byte("export const Foo = 1"), 0o644); err != nil {
		t.Fatal(err)
	}
	arts := t.TempDir()
	if err := copyFrontendComponents(specs2, arts); err != nil {
		t.Fatalf("real copy err: %v", err)
	}
	if _, err := os.Stat(filepath.Join(arts, "frontend", "src", "components", "Foo.tsx")); err != nil {
		t.Errorf("expected Foo.tsx copied: %v", err)
	}
}

func TestCopyOPARego_ZeroCov(t *testing.T) {
	// missing policy dir → skip
	fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
	if err := copyOPARego(fs, t.TempDir()); err != nil {
		t.Fatalf("missing policy err: %v", err)
	}
	// real policy dir with .rego + a non-rego file (skipped) + subdir (skipped)
	specs := t.TempDir()
	pol := filepath.Join(specs, "policy")
	if err := os.MkdirAll(filepath.Join(pol, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pol, "auth.rego"), []byte("package auth"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pol, "readme.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	arts := t.TempDir()
	if err := copyOPARego(&yongol.Fullstack{SpecsDir: specs}, arts); err != nil {
		t.Fatalf("real copy err: %v", err)
	}
	if _, err := os.Stat(filepath.Join(arts, "backend", "policy", "auth.rego")); err != nil {
		t.Errorf("expected auth.rego copied: %v", err)
	}
}
