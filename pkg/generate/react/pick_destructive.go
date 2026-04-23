//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickDestructive — FrontendTheme.Destructive 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickDestructive(t *manifest.FrontendTheme) string { return t.Destructive }
