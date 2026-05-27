package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_Subscribe(t *testing.T) {
	plan := &ir.ServicePlan{TriggerKind: ir.TriggerSubscribe}
	got := renderServiceParams(plan)
	if got != "payload: any" {
		t.Errorf("subscribe params = %q, want %q", got, "payload: any")
	}
}

func TestRenderServiceParams_BodyFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "POST",
		BodyFields: []ir.BodyFieldMeta{{Name: "title"}, {Name: "content"}},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "body: any") {
		t.Errorf("POST with BodyFields should include body param, got %q", got)
	}
}

func TestRenderServiceParams_BodyFromOpsReference(t *testing.T) {
	// POST with empty BodyFields but Ops referencing LocBody
	plan := &ir.ServicePlan{
		HTTPMethod: "POST",
		BodyFields: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpPost,
				Post: &ir.PostOp{
					Args: []ir.FieldArg{
						{Key: "title", Location: ir.LocBody},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "body: any") {
		t.Errorf("POST with LocBody ops should include body param, got %q", got)
	}
}

func TestRenderServiceParams_PathFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		PathParams: []string{"id"},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "params: any") {
		t.Errorf("GET with PathParams should include params, got %q", got)
	}
}


func TestRenderServiceParams_PathFromOpsReference(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		PathParams: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpGet,
				Get: &ir.GetOp{
					Args: []ir.FieldArg{
						{Key: "id", Location: ir.LocPath},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "params: any") {
		t.Errorf("GET with LocPath ops should include params, got %q", got)
	}
}

func TestRenderServiceParams_QueryFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod:  "GET",
		QueryParams: []ir.QueryParamMeta{{Name: "page"}},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "query: any") {
		t.Errorf("GET with QueryParams should include query, got %q", got)
	}
}

func TestRenderServiceParams_QueryFromOpsReference(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		QueryParams: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpGet,
				Get: &ir.GetOp{
					Args: []ir.FieldArg{
						{Key: "status", Location: ir.LocQuery},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "query: any") {
		t.Errorf("GET with LocQuery ops should include query, got %q", got)
	}
}

func TestRenderServiceParams_NoParams(t *testing.T) {
	plan := &ir.ServicePlan{HTTPMethod: "GET"}
	got := renderServiceParams(plan)
	if got != "user?: any" {
		t.Errorf("GET with no params = %q, want %q", got, "user?: any")
	}
}

func TestRenderServiceParams_AllThree(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod:  "PUT",
		PathParams:  []string{"id"},
		BodyFields:  []ir.BodyFieldMeta{{Name: "title"}},
		QueryParams: []ir.QueryParamMeta{{Name: "v"}},
	}
	got := renderServiceParams(plan)

	// Order: params, body, query, user
	paramsIdx := strings.Index(got, "params: any")
	bodyIdx := strings.Index(got, "body: any")
	queryIdx := strings.Index(got, "query: any")
	userIdx := strings.Index(got, "user?: any")

	if paramsIdx < 0 || bodyIdx < 0 || queryIdx < 0 || userIdx < 0 {
		t.Fatalf("missing param in %q", got)
	}
	if !(paramsIdx < bodyIdx && bodyIdx < queryIdx && queryIdx < userIdx) {
		t.Errorf("wrong order in %q", got)
	}
}

func TestRenderServiceParams_GETNoBody(t *testing.T) {
	// GET should not have body even with BodyFields
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		BodyFields: []ir.BodyFieldMeta{{Name: "title"}},
	}
	got := renderServiceParams(plan)
	if strings.Contains(got, "body: any") {
		t.Errorf("GET should not include body param, got %q", got)
	}
}
