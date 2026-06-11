//ff:func feature=stml-gen type=generator control=sequence
//ff:what 페이지의 JSX return 블록을 표준 main 래퍼 + h1 제목 + Children으로 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageJSX(page stmlparser.PageSpec, sb *strings.Builder, noBodyOps map[string]bool, ctx bindCtx) {
	sb.WriteString("  return (\n")
	sb.WriteString(`    <main className="mx-auto max-w-4xl px-4 py-8 space-y-6">` + "\n")
	sb.WriteString(fmt.Sprintf(`      <h1 className="text-2xl font-bold">%s</h1>`+"\n", kebabToTitle(page.Name)))

	if len(page.Children) > 0 {
		renderPageJSXWithChildren(page.Children, sb, noBodyOps, ctx)
	} else {
		renderPageJSXFallback(page, sb, noBodyOps, ctx)
	}

	sb.WriteString("    </main>\n")
	sb.WriteString("  )\n")
}
