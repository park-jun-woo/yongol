//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM06Binds_Unit — 점표기 경로 + 존재 필드(ok) 분기 직접 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM06Binds_Unit(t *testing.T) {
	item := getOp("GetUser", nil, map[string]*openapi3.SchemaRef{
		"User": stringProp(),
		"Name": stringProp(),
	})
	entry := operationEntry{method: "GET", op: item.Get}

	binds := []stml.FieldBind{
		{Name: "User.Name"},
		{Name: "Name"},
		{Name: "Missing"},
	}
	diags := tm06Binds(binds, "GetUser", "p.html", entry)
	// "User.Name" → top-level "User" exists (ok); "Name" exists (ok); "Missing" → 1 diag.
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for Missing, got %d: %+v", len(diags), diags)
	}
}
