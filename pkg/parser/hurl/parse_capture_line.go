//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what parseCaptureLine — [Captures] 한 줄을 HurlCapture 로 변환

package hurl

import "strings"

// parseCaptureLine parses a single [Captures] entry. Two forms are
// recognised: `<var>: jsonpath "$.x"` and `<var>: header "X-Y"`; any
// other source expression is stored verbatim so rule authors can
// special-case later.
func parseCaptureLine(line string, lineNum int) (HurlCapture, bool) {
	m := reCaptureLine.FindStringSubmatch(line)
	if m == nil {
		return HurlCapture{}, false
	}
	c := HurlCapture{Var: m[1], Line: lineNum}
	expr := strings.TrimSpace(m[2])
	if jp := reJSONPath.FindStringSubmatch(expr); jp != nil {
		c.Source = "jsonpath"
		c.JSONPath = jp[1]
		return c, true
	}
	if h := reHeaderSource.FindStringSubmatch(expr); h != nil {
		c.Source = "header"
		c.Header = h[1]
		return c, true
	}
	c.Source = expr
	return c, true
}
