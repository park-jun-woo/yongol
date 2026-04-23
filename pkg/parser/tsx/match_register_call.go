//ff:func feature=tsx-parser type=util control=sequence
//ff:what matchRegisterCall — register('field'[, { required: bool }]) 패턴 매칭

package tsx

import "encoding/json"

// matchRegisterCall recognizes `register('field'[, { required: true|false }])`.
// Works whether called standalone or as part of a {...register('x')} spread.
// required is best-effort from a BooleanLiteral inside the options object.
func matchRegisterCall(callee json.RawMessage, args []json.RawMessage) (string, bool, bool) {
	var ident struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(callee, &ident); err != nil {
		return "", false, false
	}
	if ident.Type != "Identifier" || ident.Value != "register" {
		return "", false, false
	}
	if len(args) == 0 {
		return "", false, false
	}
	var firstArg struct {
		Expression struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"expression"`
	}
	if err := json.Unmarshal(args[0], &firstArg); err != nil {
		return "", false, false
	}
	if firstArg.Expression.Type != "StringLiteral" || firstArg.Expression.Value == "" {
		return "", false, false
	}
	name := firstArg.Expression.Value
	required := false
	if len(args) >= 2 {
		required = parseRequiredFromOptions(args[1])
	}
	return name, required, true
}
