//ff:func feature=gen-react type=accessor control=sequence
//ff:what pickBackground — FrontendTheme.Background 색상 접근자

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

func pickBackground(t *manifest.FrontendTheme) string { return t.Background }
