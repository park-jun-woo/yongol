//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickMuted — FrontendTheme.Muted 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickMuted(t *manifest.FrontendTheme) string { return t.Muted }
