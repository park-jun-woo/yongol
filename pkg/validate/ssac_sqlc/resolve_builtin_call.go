//ff:func feature=validate type=rule control=selection topic=ssac-sqlc
//ff:what resolveBuiltinCall — @call / @publish 에서 (pkg, method) 쌍 추출

package ssac_sqlc

import ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// resolveBuiltinCall extracts the (pkg, method) pair that XQS-19 should
// check for a single SSaC sequence. Returns empty strings when the
// sequence is not a DB-facing built-in call.
//
//	@call <pkg>.<Method>   → (pkg, Method)
//	@publish "topic"       → ("queue", "Publish")
//
// @subscribe is handled one level up in the loop because it lives on the
// ServiceFunc, not the sequence.
func resolveBuiltinCall(seq ssacparser.Sequence, _ bool) (string, string) {
	switch seq.Type {
	case "call":
		if seq.Package == "" || seq.Model == "" {
			return "", ""
		}
		return seq.Package, seq.Model
	case "publish":
		return "queue", "Publish"
	}
	return "", ""
}
