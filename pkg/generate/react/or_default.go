//ff:func feature=gen-react type=util control=sequence
//ff:what orDefault — manifest theme 값이 비어있으면 기본값 반환

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// orDefault returns the manifest-supplied value when non-empty, else def.
func orDefault(theme *manifest.FrontendTheme, pick func(*manifest.FrontendTheme) string, def string) string {
	if theme == nil {
		return def
	}
	v := pick(theme)
	if v == "" {
		return def
	}
	return v
}
