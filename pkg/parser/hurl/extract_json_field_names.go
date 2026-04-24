//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what extractJSONFieldNames — JSON object body 에서 최상위 key 이름 추출

package hurl

import (
	"encoding/json"
	"strings"
)

// extractJSONFieldNames parses body as a JSON object and returns its
// top-level keys. Non-object payloads (arrays, primitives) and malformed
// JSON yield no keys — XOH-03 then silently skips the entry rather than
// reporting a spurious "field absent" error.
//
// Note: hurl accepts "template" bodies containing `{{var}}` placeholders
// which are not valid JSON until variables are substituted. To tolerate
// this, we replace `{{name}}` with the literal string `__hurl_var__`
// before parsing. The replacement preserves key positions so the top
// level key set is still extracted correctly.
func extractJSONFieldNames(body string) []string {
	body = strings.TrimSpace(body)
	if !strings.HasPrefix(body, "{") {
		return nil
	}
	cleaned := replaceHurlVars(body)
	var m map[string]json.RawMessage
	if err := json.Unmarshal([]byte(cleaned), &m); err != nil {
		return nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

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
		if i+1 < len(body) && body[i] == '{' && body[i+1] == '{' {
			end := strings.Index(body[i:], "}}")
			if end < 0 {
				b.WriteString(body[i:])
				return b.String()
			}
			// If preceded by `"` and followed by `"`, assume we are
			// inside a string literal — emit bare.
			prevQuote := b.Len() > 0 && b.String()[b.Len()-1] == '"'
			nextOff := i + end + 2
			nextQuote := nextOff < len(body) && body[nextOff] == '"'
			if prevQuote && nextQuote {
				b.WriteString("__hurl_var__")
			} else {
				b.WriteString(`"__hurl_var__"`)
			}
			i = nextOff
			continue
		}
		b.WriteByte(body[i])
		i++
	}
	return b.String()
}
