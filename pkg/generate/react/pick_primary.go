//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickPrimary — FrontendTheme.Primary 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickPrimary(t *manifest.FrontendTheme) string { return t.Primary }
