//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderService — ServicePlan → FastAPI service Python 함수 소스 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderService takes a ServicePlan and produces a Python service function
// body (without imports). Imports are written separately by
// WriteFeatureImports to avoid per-function duplication.
func RenderService(plan *ir.ServicePlan, reg ir.TypeRegistry) (string, error) {
	if plan == nil {
		return "", fmt.Errorf("RenderService: nil plan")
	}
	_ = reg // reserved for future type conversion use

	var b strings.Builder
	writeServiceFunc(&b, plan)
	return b.String(), nil
}
