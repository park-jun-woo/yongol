//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what hasCSRFHeader — headers 에 X-CSRF-Token 헤더 존재 여부

package hurl_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// hasCSRFHeader returns true when any header is named X-CSRF-Token
// (case-insensitive).
func hasCSRFHeader(headers []hurl.HurlHeader) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Name, "X-CSRF-Token") {
			return true
		}
	}
	return false
}
