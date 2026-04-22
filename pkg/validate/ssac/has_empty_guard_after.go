//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what hasEmptyGuardAfter — 이후 시퀀스에 @empty/@exists 가드가 있는지 판정

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// hasEmptyGuardAfter reports whether a subsequent @empty or @exists guards
// the given variable. Both guards constitute a conscious nil-check by the
// developer — @empty for "must exist" (404) and @exists for "must not exist"
// (409) — so either satisfies the FK reference guard requirement (BUG005 fix).
func hasEmptyGuardAfter(remaining []parsessac.Sequence, varName string) bool {
	for _, seq := range remaining {
		if (seq.Type == "empty" || seq.Type == "exists") && seq.Target == varName {
			return true
		}
	}
	return false
}
