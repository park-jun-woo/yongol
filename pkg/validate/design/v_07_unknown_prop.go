//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-07 — 미지의 component property (spec 정의 외) WARNING 검증
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// knownComponentProps lists the well-known component property names.
// base, variants, sizes, defaultVariant, defaultSize are ComponentToken
// struct fields (YAML-deserialized outside Props), but are listed here
// defensively in case a user accidentally nests them under props:.
var knownComponentProps = map[string]bool{
	"variant":        true,
	"size":           true,
	"color":          true,
	"disabled":       true,
	"fullWidth":      true,
	"icon":           true,
	"label":          true,
	"children":       true,
	"onClick":        true,
	"onChange":       true,
	"value":          true,
	"placeholder":    true,
	"type":           true,
	"name":           true,
	"required":       true,
	"className":      true,
	"style":          true,
	"base":           true,
	"variants":       true,
	"sizes":          true,
	"defaultVariant": true,
	"defaultSize":    true,
}

// v07UnknownProp warns about component properties not in the known set.
func v07UnknownProp(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for compName, comp := range fs.DesignSpec.Components {
		diags = append(diags, checkUnknownProps(fs.DesignSpec.File, compName, comp.Props)...)
	}
	return diags
}
