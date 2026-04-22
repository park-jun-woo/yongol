//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what itoa — strconv.Itoa 의 얇은 래퍼 (fmt import 절감)

package ssac_func

import "strconv"

// itoa is a thin wrapper to avoid fmt import bloat.
func itoa(n int) string { return strconv.Itoa(n) }
