//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickSecondary — FrontendTheme.Secondary 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickSecondary(t *manifest.FrontendTheme) string { return t.Secondary }
