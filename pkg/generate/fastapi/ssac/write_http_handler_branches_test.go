//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHTTPHandlerBranches — writeHTTPHandler 미커버 분기(필수 쿼리/pre-auth/full)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerBranches(t *testing.T) {
	t.Run("RequiredQueryAndPath", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "list_items",
			HTTPMethod:  "GET",
			URLPath:     "/item/:org_id",
			PathParams:  []string{"org_id"},
			QueryParams: []ir.QueryParamMeta{
				{Name: "status", Type: "string", Required: true},
				{Name: "page", Type: "integer", Required: false},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		out := b.String()
		if !strings.Contains(out, "    status: str,\n") {
			t.Errorf("expected required query param, got:\n%s", out)
		}
		if !strings.Contains(out, "    page: int | None = None,\n") {
			t.Errorf("expected optional query param, got:\n%s", out)
		}
		if !strings.Contains(out, "svc.list_items(session, org_id, status, page, current_user)") {
			t.Errorf("unexpected call args:\n%s", out)
		}
	})

	t.Run("PreAuthSkipsCurrentUser", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "login",
			HTTPMethod:  "POST",
			URLPath:     "/auth/login",
			BodyFields:  []ir.BodyFieldMeta{{}},
			Ops:         []ir.Op{{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{Model: "User"}}},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		out := b.String()
		if strings.Contains(out, "get_current_user") {
			t.Errorf("pre-auth handler should skip current_user, got:\n%s", out)
		}
		if !strings.Contains(out, "    body: LoginRequest,\n") {
			t.Errorf("expected body param, got:\n%s", out)
		}
		if !strings.Contains(out, "svc.login(session, body)") {
			t.Errorf("unexpected call args:\n%s", out)
		}
	})

	t.Run("PublishAddsEventBus", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "do_pub",
			HTTPMethod:  "POST",
			URLPath:     "/x/pub",
			Ops:         []ir.Op{{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "t"}}},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		out := b.String()
		if !strings.Contains(out, "event_bus: EventBus = Depends(get_event_bus)") {
			t.Errorf("expected event_bus dep, got:\n%s", out)
		}
		if !strings.Contains(out, "event_bus)") {
			t.Errorf("expected event_bus call arg, got:\n%s", out)
		}
	})
}

func TestOpenAPITypeToPython(t *testing.T) {
	cases := map[string]string{
		"integer": "int",
		"number":  "float",
		"boolean": "bool",
		"string":  "str",
		"unknown": "str",
	}
	for in, want := range cases {
		if got := openAPITypeToPython(in); got != want {
			t.Errorf("openAPITypeToPython(%q) = %q, want %q", in, got, want)
		}
	}
}
