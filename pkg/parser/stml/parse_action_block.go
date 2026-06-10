//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what data-action 요소에서 ActionBlock 구성
package stml

import (
	"strings"

	"golang.org/x/net/html"
)

// parseActionBlock builds an ActionBlock from a data-action element.
func parseActionBlock(n *html.Node, operationID string) ActionBlock {
	ab := ActionBlock{
		Tag:         n.Data,
		ClassName:   getAttr(n, "class"),
		OperationID: operationID,
		Params:      extractParams(n),
		EnabledWhen: getAttr(n, "data-enabled-when"),
		Invalidates: strings.Fields(getAttr(n, "data-invalidates")),
		CaptureRaw:  getAttr(n, "data-capture"),
		Redirect:    getAttr(n, "data-redirect"),
		OnErrorNode: hasDescendantOnError(n),
	}
	if ab.CaptureRaw != "" {
		// Syntax errors are surfaced by TM-20 at validate time (re-parse).
		ab.Captures, _ = ParseCapture(ab.CaptureRaw)
	}
	if n.Data == "button" {
		ab.SubmitText = directText(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkActionChildren(c, &ab)
	}
	return ab
}
