//ff:func feature=validate type=rule control=sequence topic=design-structural
//ff:what V-07 — 미지의 component property (spec 정의 외) WARNING 검증
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// knownComponentProps lists the well-known component property names.
var knownComponentProps = map[string]bool{
	"variant":   true,
	"size":      true,
	"color":     true,
	"disabled":  true,
	"fullWidth": true,
	"icon":      true,
	"label":     true,
	"children":  true,
	"onClick":   true,
	"onChange":   true,
	"value":     true,
	"placeholder": true,
	"type":      true,
	"name":      true,
	"required":  true,
	"className": true,
	"style":     true,
}

// v07UnknownProp warns about component properties not in the known set.
func v07UnknownProp(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for compName, comp := range fs.DesignSpec.Components {
		for propName := range comp.Props {
			if !knownComponentProps[propName] {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fs.DesignSpec.File,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: "[V-07] component \"" + compName + "\" has unknown property: \"" + propName + "\"",
					Advice:  "Verify this property name is intentional; known props: variant, size, color, disabled, fullWidth, icon, label, children, onClick, onChange, value, placeholder, type, name, required, className, style",
				})
			}
		}
	}
	return diags
}
