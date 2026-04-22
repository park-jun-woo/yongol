//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-openapi
//ff:what XOS-66 — SSaC가 참조한 request 필드가 OpenAPI requestBody required에 포함되어 있는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos66UsedFieldsRequired validates XOS-66: every SSaC-referenced
// `request.<field>` is declared in the OpenAPI requestBody's required list.
func xos66UsedFieldsRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		rs, ok := fs.RequestConstraints[fn.Name]
		if !ok {
			continue
		}
		for field := range collectRequestFields(fn) {
			fc, exists := rs[field]
			if !exists || fc.Required {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    fn.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XOS-66] field " + field + " is used in SSaC " + fn.Name + " but not marked required in OpenAPI requestBody",
				Advice:  "SSaC 가 사용하는 필드 " + field + " 를 OpenAPI requestBody 의 required 배열에 추가하세요",
			})
		}
	}
	return diags
}
