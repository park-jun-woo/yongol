//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestXMO12_NoFrontButConsumed_Warning — no-front 태그인데 STML이 소비 중 → WARNING

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO12_NoFrontButConsumed_Warning(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
	}}
	noFrontOp := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "ListUsers", Tags: []string{"no-front"}}}
	doc := makeDoc(map[string]*openapi3.PathItem{"/users": noFrontOp})
	diags := Run(makeFS(pages, doc))
	if countDiag(diags, "[XMO-12]") != 1 {
		t.Fatalf("expected 1 XMO-12, got %d (%v)", countDiag(diags, "[XMO-12]"), diags)
	}
	for _, d := range diags {
		if strings.HasPrefix(d.Message, "[XMO-12]") && d.Level != diagnostic.LevelWarning {
			t.Errorf("XMO-12 level = %q, want WARNING", d.Level)
		}
	}
}
