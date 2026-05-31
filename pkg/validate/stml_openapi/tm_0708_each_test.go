//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestTM0708Each — TM-07(미존재)/TM-08(비배열)/배열(ok) 분기 검증
package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM0708Each(t *testing.T) {
	arrayProp := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}}
	item := getOp("ListItems", nil, map[string]*openapi3.SchemaRef{
		"items": arrayProp,
		"count": stringProp(),
	})
	entry := operationEntry{method: "GET", op: item.Get}

	eaches := []stml.EachBlock{
		{Field: "items"},   // array → ok
		{Field: "count"},   // not array → TM-08
		{Field: "missing"}, // not in schema → TM-07
	}
	diags := tm0708Each(eaches, "ListItems", "p.html", entry)
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(diags), diags)
	}
	var has07, has08 bool
	for _, d := range diags {
		if strings.Contains(d.Message, "[TM-07]") {
			has07 = true
		}
		if strings.Contains(d.Message, "[TM-08]") {
			has08 = true
		}
	}
	if !has07 || !has08 {
		t.Fatalf("expected both TM-07 and TM-08, got %+v", diags)
	}
}
