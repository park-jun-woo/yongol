//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderFeatureSchemas — feature 내 body 필드 보유 plan 들의 Pydantic BaseModel 클래스 렌더

package fastapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderFeatureSchemas produces Pydantic BaseModel classes for all plans in
// a feature that have request body fields. Returns empty string when no
// schemas are needed.
func renderFeatureSchemas(plans []*ir.ServicePlan) string {
	var b strings.Builder
	hasSchema := false

	for _, plan := range plans {
		if !planHasRequestBody(plan) {
			continue
		}
		if !hasSchema {
			b.WriteString("from pydantic import BaseModel\n")
			b.WriteString("from typing import Optional\n\n\n")
			hasSchema = true
		}
		renderOneSchema(&b, plan)
	}

	if !hasSchema {
		return ""
	}
	return b.String()
}
