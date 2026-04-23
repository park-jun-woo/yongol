//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickForeground — FrontendTheme.Foreground 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickForeground(t *manifest.FrontendTheme) string { return t.Foreground }
