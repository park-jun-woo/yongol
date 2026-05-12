//ff:func feature=gen-react type=util control=sequence
//ff:what DESIGN.md 색상값을 키로 조회하고 없으면 기본값을 반환한다

package react

import "github.com/park-jun-woo/yongol/pkg/parser/design"

// designColor returns the DESIGN.md color value for a key, falling back to def.
func designColor(dspec *design.DesignSpec, key, def string) string {
	if dspec == nil || dspec.Colors == nil {
		return def
	}
	if v, ok := dspec.Colors[key]; ok {
		return v
	}
	return def
}
