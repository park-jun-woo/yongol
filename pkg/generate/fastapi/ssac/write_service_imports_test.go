//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteServiceImports — feature service 파일 통합 import 작성
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteServiceImports(t *testing.T) {
	plans := []*ir.ServicePlan{
		{
			OperationID: "delete_item",
			Feature:     "billing",
			Ops: []ir.Op{
				{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "item"}},
				{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "item.deleted"}},
			},
		},
	}
	var b strings.Builder
	writeServiceImports(&b, plans, "billing")
	got := b.String()
	for _, want := range []string{
		"from fastapi import HTTPException",
		"from sqlalchemy import delete",
		"from sqlalchemy.ext.asyncio import AsyncSession",
		"from app.models.models import Item",
		"from app.dependencies.event_bus import EventBus",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	// Should end with a trailing blank line.
	if !strings.HasSuffix(got, "\n\n") {
		t.Errorf("expected trailing blank line, got:\n%q", got)
	}
}
