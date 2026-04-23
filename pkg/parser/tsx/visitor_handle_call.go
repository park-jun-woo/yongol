//ff:func feature=tsx-parser type=parser control=sequence
//ff:what (v *visitor).handleCall — apiClient.<op>(...) / register('name', ...) 두 패턴 동시 추출

package tsx

import "encoding/json"

// handleCall extracts apiClient.<op>(...) and register('name', ...) calls.
// A single CallExpression can match neither, one, or both (rare) patterns;
// each path returns independently.
func (v *visitor) handleCall(raw json.RawMessage) {
	var c struct {
		Span      astSpan           `json:"span"`
		Callee    json.RawMessage   `json:"callee"`
		Arguments []json.RawMessage `json:"arguments"`
	}
	if err := json.Unmarshal(raw, &c); err != nil {
		return
	}

	// Pattern 1: apiClient.<op>(...)
	if opID, ok := matchApiClientCallee(c.Callee); ok {
		line, col := v.resolve(calleePropertySpan(c.Callee).Start)
		call := APICall{OperationID: opID, Kind: "raw", Line: line, Col: col}
		if len(c.Arguments) > 0 {
			call.Args = v.extractArgBindings(c.Arguments[0])
		}
		v.page.Calls = append(v.page.Calls, call)
	}

	// Pattern 2: register('name', opts?)
	if name, required, ok := matchRegisterCall(c.Callee, c.Arguments); ok {
		line, col := v.resolve(c.Span.Start)
		v.page.FormFields = append(v.page.FormFields, FormField{
			Name: name, Required: required, Line: line, Col: col,
		})
	}
}
