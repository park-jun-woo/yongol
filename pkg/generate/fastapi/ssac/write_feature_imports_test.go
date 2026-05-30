//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteFeatureImports — feature 단위 통합 import 블록 생성

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteFeatureImports(t *testing.T) {
	t.Run("Minimal", func(t *testing.T) {
		got := WriteFeatureImports(nil, "billing")
		if !strings.Contains(got, "from fastapi import HTTPException") {
			t.Errorf("expected base imports, got:\n%s", got)
		}
		if !strings.Contains(got, "from sqlalchemy.ext.asyncio import AsyncSession") {
			t.Errorf("expected AsyncSession import, got:\n%s", got)
		}
	})
	t.Run("WithModelAndSelect", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "get_item",
				Feature:     "billing",
				Ops: []ir.Op{
					{Kind: ir.OpGet, Get: &ir.GetOp{Model: "item"}},
				},
			},
		}
		got := WriteFeatureImports(plans, "billing")
		if !strings.Contains(got, "from sqlalchemy import select") {
			t.Errorf("expected select import, got:\n%s", got)
		}
		if !strings.Contains(got, "from app.models.models import Item") {
			t.Errorf("expected model import, got:\n%s", got)
		}
	})
}
