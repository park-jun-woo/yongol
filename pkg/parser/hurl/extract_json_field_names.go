//ff:func feature=crosscheck type=util control=iteration dimension=1 topic=scenario-check
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
