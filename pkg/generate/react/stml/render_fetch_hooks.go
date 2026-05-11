//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what FetchBlock의 useQuery 훅 선언을 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFetchHooks writes useQuery hook declarations.
func renderFetchHooks(f stmlparser.FetchBlock, sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("  %s\n\n", renderUseQuery(f)))
	for _, child := range f.NestedFetches {
		renderFetchHooks(child, sb)
	}
}
