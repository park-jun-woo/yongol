//ff:func feature=validate type=util control=iteration dimension=1 topic=design-structural
//ff:what checkUnknownProps — 단일 component의 props 중 미지 항목에 WARNING 진단 생성
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkUnknownProps returns diagnostics for unknown property names in a component.
func checkUnknownProps(file, compName string, props map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for propName := range props {
		if !knownComponentProps[propName] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[V-07] component \"" + compName + "\" has unknown property: \"" + propName + "\"",
				Advice:  "Verify this property name is intentional; known props: variant, size, color, disabled, fullWidth, icon, label, children, onClick, onChange, value, placeholder, type, name, required, className, style",
			})
		}
	}
	return diags
}
