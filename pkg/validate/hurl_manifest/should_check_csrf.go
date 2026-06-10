//ff:func feature=validate type=rule control=sequence topic=hurl-manifest
//ff:what shouldCheckCSRF — entry 가 CSRF 검사 대상인지 판정 (mutating + 비-auth + headerName 미포함)

package hurl_manifest

import "github.com/park-jun-woo/yongol/pkg/parser/hurl"

// shouldCheckCSRF reports whether e is a mutating non-auth request
// lacking the manifest-resolved CSRF header (headerName). When true the
// caller emits an XOH-07 WARNING for it.
func shouldCheckCSRF(e hurl.HurlEntry, headerName string) bool {
	if !isMutating(e.Method) {
		return false
	}
	if isAuthPath(e.Path) {
		return false
	}
	if hasCSRFHeader(e.Headers, headerName) {
		return false
	}
	return true
}
