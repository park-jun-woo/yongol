//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what DOM 전체를 순회하며 data-link·data-action 동시 선언(클릭 의미 충돌)을 파스 진단으로 수집
package stml

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

// collectLinkActionConflicts walks the whole DOM and reports every element
// declaring both data-link and data-action: one element cannot both
// navigate and submit on click, so the conflict is rejected at parse time
// (page-flow Phase007).
func collectLinkActionConflicts(n *html.Node, file string, out *[]diagnostic.Diagnostic) {
	if n.Type == html.ElementNode && getAttr(n, "data-link") != "" && getAttr(n, "data-action") != "" {
		*out = append(*out, diagnostic.Diagnostic{
			File:    file,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("data-link %q and data-action %q declared on the same element — click semantics conflict (navigate vs submit)", getAttr(n, "data-link"), getAttr(n, "data-action")),
			Advice:  "Split the link and the action into separate elements",
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectLinkActionConflicts(c, file, out)
	}
}
