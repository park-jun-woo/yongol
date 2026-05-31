//ff:func feature=gen-ir type=test control=sequence
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPublish(t *testing.T) {
	op := convertPublish(ssac.Sequence{
		Topic:   "order.completed",
		Inputs:  map[string]string{"ID": "order.ID"},
		Options: map[string]string{"delay": "1800"},
	})
	if op.Kind != OpPublish || op.Publish == nil {
		t.Fatalf("op = %+v", op)
	}
	if op.Publish.Topic != "order.completed" {
		t.Errorf("topic = %q", op.Publish.Topic)
	}
	if len(op.Publish.Payload) != 1 || len(op.Publish.Options) != 1 {
		t.Errorf("payload=%v options=%v", op.Publish.Payload, op.Publish.Options)
	}
}
