//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickRadius — FrontendTheme.Radius 값 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickRadius(t *manifest.FrontendTheme) string { return t.Radius }
