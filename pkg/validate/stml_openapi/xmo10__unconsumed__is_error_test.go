//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestXMO10_Unconsumed_IsError — Frontend ON 미소비 op은 ERROR 레벨로 승격

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_IsError(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListUsers"}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": getOp("GetUser", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if countDiag(diags, "[XMO-10]") != 1 {
		t.Fatalf("expected 1 XMO-10, got %d (%v)", countDiag(diags, "[XMO-10]"), diags)
	}
	for _, d := range diags {
		if strings.HasPrefix(d.Message, "[XMO-10]") && d.Level != diagnostic.LevelError {
			t.Errorf("XMO-10 level = %q, want ERROR", d.Level)
		}
	}
}
