//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 페이지의 useParams, useQueryClient, useQuery 훅을 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageHooks(page stmlparser.PageSpec, is importSet, sb *strings.Builder) {
	allParams := collectAllParams(page)
	if up := renderUseParams(allParams); up != "" {
		sb.WriteString(fmt.Sprintf("  %s\n", up))
	}

	if is.useQueryClient {
		sb.WriteString("  const queryClient = useQueryClient()\n")
	}

	sb.WriteString("\n")

	for _, f := range page.Fetches {
		renderFetchHooks(f, sb)
	}
}
