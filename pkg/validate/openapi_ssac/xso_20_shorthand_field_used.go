//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XSO-20 — OpenAPI response 필드가 shorthand @response 변수 타입에서 사용되는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xso20ShorthandFieldUsed validates XSO-20: every OpenAPI response property is
// referenced by the shorthand @response variable type. Page[T]/Cursor[T]/[]T
// wrappers are skipped (defeater).
func xso20ShorthandFieldUsed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		varName := shorthandResponseTarget(fn)
		if varName == "" {
			continue
		}
		// Page[T]/Cursor[T] wrapper 면제 제거 — B안에서 shorthand @response도
		// OpenAPI response fields와 매칭 검증을 받아야 함.
		fields := g.Schemas["SSaC.response."+fn.Name]
		if len(fields) == 0 {
			continue
		}
		opProps := g.Schemas["OpenAPI.response."+fn.Name]
		if len(opProps) == 0 {
			continue
		}
		diags = append(diags, collectUnusedProps(fn, opProps, toSet(fields), "XSO-20")...)
	}
	return diags
}
