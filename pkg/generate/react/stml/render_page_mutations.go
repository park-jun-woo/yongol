//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 페이지의 모든 Action에 대한 useForm + useMutation 훅을 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageMutations(allActions []stmlparser.ActionBlock, fetchOps []string, sb *strings.Builder) {
	for _, a := range allActions {
		if len(a.Fields) > 0 {
			sb.WriteString(fmt.Sprintf("  %s\n", renderFormHook(a)))
		}
		sb.WriteString(fmt.Sprintf("  %s\n\n", renderUseMutation(a, fetchOps)))
	}
}
