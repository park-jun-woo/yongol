//ff:func feature=tsx-parser type=util control=sequence
//ff:what matchApiClientCallee — MemberExpression(object=apiClient, property=<op>) 매칭 → operationId 반환

package tsx

import "encoding/json"

// matchApiClientCallee checks that the callee is MemberExpression with
// object=Identifier("apiClient") and property=Identifier(<op>). Returns
// the operationId and true on success. apiClient.foo.bar() is rejected.
func matchApiClientCallee(callee json.RawMessage) (string, bool) {
	var m struct {
		Type   string `json:"type"`
		Object struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"object"`
		Property struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"property"`
	}
	if err := json.Unmarshal(callee, &m); err != nil {
		return "", false
	}
	if m.Type != "MemberExpression" {
		return "", false
	}
	if m.Object.Type != "Identifier" || m.Object.Value != "apiClient" {
		return "", false
	}
	if m.Property.Type != "Identifier" || m.Property.Value == "" {
		return "", false
	}
	return m.Property.Value, true
}
