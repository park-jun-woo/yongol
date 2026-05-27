//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteServiceFuncSubscribe — Subscribe handler message = payload alias 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteServiceFuncSubscribe(t *testing.T) {
	t.Run("MessageAlias", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "OnWorkflowExecuted",
			TriggerKind: ir.TriggerSubscribe,
			Topic:       "workflow.executed",
			Ops:         []ir.Op{},
		}
		var b strings.Builder
		writeServiceFunc(&b, plan)
		got := b.String()
		if !strings.Contains(got, "message = payload") {
			t.Errorf("expected message = payload alias, got: %s", got)
		}
	})

	t.Run("HTTPNoMessageAlias", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetWorkflow",
			TriggerKind: ir.TriggerHTTP,
			HTTPMethod:  "GET",
			Ops:         []ir.Op{},
		}
		var b strings.Builder
		writeServiceFunc(&b, plan)
		got := b.String()
		if strings.Contains(got, "message = payload") {
			t.Errorf("HTTP handler should not have message alias, got: %s", got)
		}
	})
}
