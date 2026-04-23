//ff:func feature=cli type=test-helper control=sequence
//ff:what isUsageError — err 이 usageArgs / FlagErrorFunc 경유인지 판별

package main

import "errors"

// isUsageError reports whether err originated from usageArgs / the root
// FlagErrorFunc. Matches main.go's exit-code 2 branch.
func isUsageError(err error) bool {
	var ue *usageError
	return errors.As(err, &ue)
}
