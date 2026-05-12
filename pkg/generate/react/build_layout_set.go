//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what LayoutSpec 목록에서 레이아웃 이름 set을 생성한다

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// buildLayoutSet returns a set of layout names that have corresponding LayoutSpecs.
func buildLayoutSet(layouts []stml.LayoutSpec) map[string]bool {
	s := make(map[string]bool, len(layouts))
	for _, l := range layouts {
		s[l.Name] = true
	}
	return s
}
