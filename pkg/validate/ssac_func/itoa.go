//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what itoa — thin wrapper around strconv.Itoa to avoid importing fmt

package ssac_func

import "strconv"

// itoa is a thin wrapper to avoid fmt import bloat.
func itoa(n int) string { return strconv.Itoa(n) }
