//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestCollectSchemaModels — request body 있는 plan에서 Pydantic 모델명 수집/중복제거

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectSchemaModels(t *testing.T) {
	body := []ir.BodyFieldMeta{{}}

	t.Run("Empty", func(t *testing.T) {
		if got := collectSchemaModels(nil); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
	t.Run("GetSkipped", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "list_items", HTTPMethod: "GET", BodyFields: body},
		}
		if got := collectSchemaModels(plans); got != nil {
			t.Errorf("GET should be skipped, got %v", got)
		}
	})
	t.Run("PostWithoutBodySkipped", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "create_item", HTTPMethod: "POST"},
		}
		if got := collectSchemaModels(plans); got != nil {
			t.Errorf("no body should be skipped, got %v", got)
		}
	})
	t.Run("PostPutPatchWithBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "create_item", HTTPMethod: "POST", BodyFields: body},
			{OperationID: "update_item", HTTPMethod: "PUT", BodyFields: body},
			{OperationID: "patch_item", HTTPMethod: "PATCH", BodyFields: body},
		}
		got := collectSchemaModels(plans)
		want := []string{"CreateItemRequest", "UpdateItemRequest", "PatchItemRequest"}
		if len(got) != 3 {
			t.Fatalf("got %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got[%d]=%q, want %q", i, got[i], want[i])
			}
		}
	})
	t.Run("Dedup", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "create_item", HTTPMethod: "POST", BodyFields: body},
			{OperationID: "create_item", HTTPMethod: "post", BodyFields: body},
		}
		got := collectSchemaModels(plans)
		if len(got) != 1 || got[0] != "CreateItemRequest" {
			t.Errorf("expected dedup to single entry, got %v", got)
		}
	})
}
