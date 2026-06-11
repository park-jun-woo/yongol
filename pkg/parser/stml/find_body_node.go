//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what findBodyNode — html.Parse 결과 트리에서 <body> 요소 노드 탐색

package stml

import "golang.org/x/net/html"

// findBodyNode returns the first <body> element in the parsed tree, or nil.
// html.Parse always synthesizes html/head/body, so fragment input still
// lands its top-level elements under the returned node.
func findBodyNode(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if body := findBodyNode(c); body != nil {
			return body
		}
	}
	return nil
}
