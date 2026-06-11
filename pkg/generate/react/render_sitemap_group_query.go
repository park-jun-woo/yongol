//ff:func feature=gen-react type=generator control=sequence
//ff:what renderSitemapGroupQuery — 동적 그룹 1개 op 의 레이아웃 useQuery const 방출 (페이지와 동일한 쿼리 키 규약, bearer 는 token 게이트)

package react

import (
	"fmt"
	"strings"
)

// renderSitemapGroupQuery writes the layout useQuery const for one dynamic
// menu group operation (plans/stml/sitemap Phase007). The query key is the
// page fetch convention exactly — ['<OperationID>'] (renderUseQuery's key
// for a parameterless fetch) — so a page action's data-invalidates on the
// same operation hits the menu query with no new vocabulary. In bearer
// mode the query is gated with enabled: !!token (the layout-level
// useAuthStore token selector) so a signed-out visitor never fires the
// protected call; cookie mode has no gate — the client cannot know the
// session state, and an error response simply renders no group (the
// renderer's zero-item omission).
func renderSitemapGroupQuery(sb *strings.Builder, op string, bearerGate bool) {
	fmt.Fprintf(sb, "  const { data: %sData } = useQuery({\n", lowerFirst(op))
	fmt.Fprintf(sb, "    queryKey: ['%s'],\n", op)
	fmt.Fprintf(sb, "    queryFn: () => api.%s(),\n", op)
	if bearerGate {
		sb.WriteString("    enabled: !!token,\n")
	}
	sb.WriteString("  })\n")
}
