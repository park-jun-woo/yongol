//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what replaceHurlVarAt — body[i] 위치에서 `{{...}}` placeholder 를 치환 (해당되지 않으면 no-op)

package hurl

import "strings"

// replaceHurlVarAt inspects body at offset i. When it starts a `{{...}}`
// hurl placeholder the helper writes the appropriate replacement into b
// and returns the new offset (past the closing `}}`). When no placeholder
// matches the helper returns (i, false) and the caller appends body[i]
// verbatim. When the placeholder is unterminated (no `}}` found), the
// helper appends body[i:] and returns (len(body), true).
func replaceHurlVarAt(body string, i int, b *strings.Builder) (int, bool) {
	if i+1 >= len(body) || body[i] != '{' || body[i+1] != '{' {
		return i, false
	}
	end := strings.Index(body[i:], "}}")
	if end < 0 {
		b.WriteString(body[i:])
		return len(body), true
	}
	// If preceded by `"` and followed by `"`, assume we are inside a
	// string literal — emit bare.
	prevQuote := b.Len() > 0 && b.String()[b.Len()-1] == '"'
	nextOff := i + end + 2
	nextQuote := nextOff < len(body) && body[nextOff] == '"'
	if prevQuote && nextQuote {
		b.WriteString("__hurl_var__")
	} else {
		b.WriteString(`"__hurl_var__"`)
	}
	return nextOff, true
}
