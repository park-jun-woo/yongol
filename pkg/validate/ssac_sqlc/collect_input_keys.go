//ff:func feature=validate type=util control=iteration dimension=1 topic=sqlc
//ff:what collectInputKeys — sequence의 Args.Field 와 Inputs key를 모아 반환

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// collectInputKeys returns the input key names for a sequence.
// CRUD sequences carry inputs in Args (field names); state/auth/publish
// sequences carry them in the Inputs map (snake_case keys).
func collectInputKeys(seq ssac.Sequence) []string {
	var keys []string
	for _, arg := range seq.Args {
		if arg.Field == "" {
			continue
		}
		keys = append(keys, arg.Field)
	}
	for k := range seq.Inputs {
		if k == "" {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}
