//ff:func feature=crosscheck type=util control=iteration dimension=1 topic=scenario-check
//ff:what replaceHurlVars — `{{var}}` 를 `__hurl_var__` 로 치환해 JSON decode 가능하도록 정규화

package hurl

import "strings"

// replaceHurlVars replaces every `{{var}}` with the literal token
// `__hurl_var__` so that JSON decoding succeeds when the body still
// carries hurl template placeholders. The replacement text is a plain
// identifier, which is why we wrap it in quotes whenever the placeholder
// stood bare as a value — otherwise the surrounding quotes that were
// already present in the source (e.g. `"{{email}}"`) give us a valid
// JSON string. Edge cases (placeholder as an object key, placeholder
// inside a number literal) remain unsupported; XOH-03 simply skips
// unparseable bodies rather than false-reporting.
func replaceHurlVars(body string) string {
	var b strings.Builder
	i := 0
	for i < len(body) {
		next, consumed := replaceHurlVarAt(body, i, &b)
		if next > i {
			i = next
			continue
		}
		if !consumed {
			b.WriteByte(body[i])
		}
		i++
	}
	return b.String()
}
