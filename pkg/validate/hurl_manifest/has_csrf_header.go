//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what hasCSRFHeader — headers 에 headerName CSRF 헤더 존재 여부 (대소문자 무관)

package hurl_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// hasCSRFHeader returns true when any header is named headerName
// (case-insensitive — HTTP header names are case-insensitive).
func hasCSRFHeader(headers []hurl.HurlHeader, headerName string) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Name, headerName) {
			return true
		}
	}
	return false
}
