//ff:func feature=gen-react type=test control=sequence
//ff:what findRefreshOp — refresh_op operationId 매칭 시 plan 구성, 미매칭 시 found=false 검증

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestFindRefreshOp(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{
			Post: buildTokenOp("Refresh", []string{"access_token", "refresh_token"}, []string{"refresh_token"}),
		}),
	)}
	fa := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token", RefreshOp: "Refresh"}

	rp, found := findRefreshOp(doc, fa)
	if !found || rp == nil {
		t.Fatalf("found=%v rp=%v, want plan for Refresh", found, rp)
	}
	if rp.opID != "Refresh" || rp.method != "POST" || rp.path != "/auth/refresh" {
		t.Errorf("plan identity = %+v", rp)
	}
	if rp.tokenField != "access_token" || rp.refreshField != "refresh_token" || rp.bodyKey != "refresh_token" {
		t.Errorf("plan fields = %+v", rp)
	}

	// operationId not present in any operation -> found=false
	missing := &manifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token", RefreshOp: "DoesNotExist"}
	if rp, found := findRefreshOp(doc, missing); found || rp != nil {
		t.Errorf("missing op = (%v,%v), want (nil,false)", rp, found)
	}
}
