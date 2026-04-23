//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickAccent — FrontendTheme.Accent 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickAccent(t *manifest.FrontendTheme) string { return t.Accent }
