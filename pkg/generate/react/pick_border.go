//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickBorder — FrontendTheme.Border 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickBorder(t *manifest.FrontendTheme) string { return t.Border }
