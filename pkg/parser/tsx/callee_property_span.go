//ff:func feature=tsx-parser type=util control=sequence
//ff:what calleePropertySpan — MemberExpression callee 의 property identifier span 반환

package tsx

import "encoding/json"

// calleePropertySpan returns the span of the property identifier in a
// MemberExpression callee. Falls back to a zero span on error so the
// caller's resolve(0) produces (0,0).
func calleePropertySpan(callee json.RawMessage) astSpan {
	var m struct {
		Property struct {
			Span astSpan `json:"span"`
		} `json:"property"`
	}
	_ = json.Unmarshal(callee, &m)
	return m.Property.Span
}
