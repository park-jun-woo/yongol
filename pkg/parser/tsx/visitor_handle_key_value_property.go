//ff:func feature=tsx-parser type=parser control=iteration dimension=1
//ff:what (v *visitor).handleKeyValueProperty — `{ mutationFn: apiClient.X }` 패턴 추출

package tsx

import "encoding/json"

// handleKeyValueProperty recognises `{ mutationFn: apiClient.X }` and
// `{ queryFn: apiClient.X }` patterns where the value is a bare member
// expression (not a CallExpression). Without this branch, TanStack Query
// style registrations slip past XOT-* rules because there is no literal
// `apiClient.X()` site for the visitor to hook onto.
func (v *visitor) handleKeyValueProperty(raw json.RawMessage) {
	var kv struct {
		Key struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"key"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(raw, &kv); err != nil {
		return
	}
	switch kv.Key.Value {
	case "mutationFn", "queryFn":
		// continue
	default:
		return
	}
	opID, ok := matchApiClientCallee(kv.Value)
	if !ok {
		return
	}
	line, col := v.resolve(calleePropertySpan(kv.Value).Start)
	// Already recorded by a concrete CallExpression? Skip to avoid doubles.
	for _, existing := range v.page.Calls {
		if existing.OperationID == opID && existing.Line == line {
			return
		}
	}
	v.page.Calls = append(v.page.Calls, APICall{
		OperationID: opID, Kind: "raw", Line: line, Col: col,
	})
}
