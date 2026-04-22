//ff:func feature=validate type=util control=iteration dimension=1 topic=sqlc
//ff:what findCaseInsensitiveParam — return a case-insensitive match from the param set

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// findCaseInsensitiveParam returns a param from the set that matches key
// case-insensitively, or "" if no such param exists.
func findCaseInsensitiveParam(key string, params rule.StringSet) string {
	lower := strings.ToLower(key)
	for p := range params {
		if strings.ToLower(p) == lower {
			return p
		}
	}
	return ""
}
