//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStatusInt(t *testing.T) {
	if got := statusInt("200"); got != 200 {
		t.Errorf("statusInt(200) = %d", got)
	}
	if got := statusInt("201"); got != 201 {
		t.Errorf("statusInt(201) = %d", got)
	}
	if got := statusInt("404"); got != 0 {
		t.Errorf("statusInt(404) = %d, want 0", got)
	}
}

func TestSchemaType(t *testing.T) {
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if got := schemaType(ref); got != "string" {
		t.Errorf("schemaType = %q, want string", got)
	}
	// nil ref / nil value / nil type → "".
	if got := schemaType(nil); got != "" {
		t.Errorf("nil ref: %q", got)
	}
	if got := schemaType(&openapi3.SchemaRef{}); got != "" {
		t.Errorf("nil value: %q", got)
	}
	if got := schemaType(&openapi3.SchemaRef{Value: &openapi3.Schema{}}); got != "" {
		t.Errorf("nil type: %q", got)
	}
	// empty type slice → "".
	if got := schemaType(&openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{}}}); got != "" {
		t.Errorf("empty type slice: %q", got)
	}
}

func TestIsAuthEndpoint(t *testing.T) {
	authOp := &openapi3.Operation{Security: &openapi3.SecurityRequirements{}}
	if !isAuthEndpoint(authOp) {
		t.Error("empty security requirement should be an auth endpoint")
	}
	// nil security → not auth.
	if isAuthEndpoint(&openapi3.Operation{}) {
		t.Error("nil security should not be auth endpoint")
	}
	// non-empty security → not auth.
	reqs := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearer": {}}}
	if isAuthEndpoint(&openapi3.Operation{Security: &reqs}) {
		t.Error("non-empty security should not be auth endpoint")
	}
}

func TestHasMatchingParam(t *testing.T) {
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "userID"}},
		},
	}
	entry := operationEntry{method: "GET", op: op}
	if !hasMatchingParam(entry, "userid") {
		t.Error("expected case-insensitive match on userID")
	}
	if hasMatchingParam(entry, "missing") {
		t.Error("missing param should not match")
	}
	// nil op → false.
	if hasMatchingParam(operationEntry{}, "x") {
		t.Error("nil op should yield no match")
	}
}

func TestDefaultLayoutFromManifest(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Frontend.DefaultLayout = "main"
	if got := defaultLayoutFromManifest(fs); got != "main" {
		t.Errorf("got %q, want main", got)
	}
	// nil manifest → "".
	if got := defaultLayoutFromManifest(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest: %q", got)
	}
}

func TestCollectPropNames(t *testing.T) {
	out := map[string]struct{}{}
	s := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{}},
			"name": {Value: &openapi3.Schema{}},
		},
		AllOf: openapi3.SchemaRefs{
			{Value: &openapi3.Schema{Properties: openapi3.Schemas{"email": {Value: &openapi3.Schema{}}}}},
			nil, // skipped
		},
	}
	collectPropNames(out, s)
	for _, want := range []string{"id", "name", "email"} {
		if _, ok := out[want]; !ok {
			t.Errorf("missing prop %q", want)
		}
	}
	// nil schema is a no-op.
	collectPropNames(out, nil)
}

func TestAddSchemaProps(t *testing.T) {
	out := map[string]responseFieldInfo{}
	s := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		},
		AllOf: openapi3.SchemaRefs{
			{Value: &openapi3.Schema{Properties: openapi3.Schemas{
				"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			}}},
		},
	}
	addSchemaProps(out, s)
	if out["id"].typ != "integer" {
		t.Errorf("id typ = %q", out["id"].typ)
	}
	if out["name"].typ != "string" {
		t.Errorf("name typ = %q", out["name"].typ)
	}
	addSchemaProps(out, nil) // no-op
}

func TestEmitIfClass(t *testing.T) {
	diags := emitIfClass("page.stml", "div", "card", "bg-red")
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[TM-10]") || !strings.Contains(diags[0].Message, "bg-red") {
		t.Errorf("unexpected message: %q", diags[0].Message)
	}
	// empty className → no diagnostic.
	if got := emitIfClass("page.stml", "div", "card", ""); got != nil {
		t.Errorf("empty class should yield nil, got %v", got)
	}
}
