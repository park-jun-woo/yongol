//ff:func feature=migration type=util control=sequence
//ff:what itoa — strconv.Itoa 재노출 (패키지 내부 유틸)
package migration

import "strconv"

// itoa is a tiny shim so other files in this package don't need to
// import strconv just for Itoa. Keeps imports tidy.
func itoa(n int) string { return strconv.Itoa(n) }
