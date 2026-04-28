//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what carriesAuthHeader — Authorization / Cookie 헤더 존재 여부 판정

package hurl_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// carriesAuthHeader returns true when at least one request header looks
// like it transports auth state — Authorization, Cookie, or the
// __Host-access_token / refresh_token cookies yongol sets by default.
func carriesAuthHeader(headers []hurl.HurlHeader) bool {
	for _, h := range headers {
		name := strings.ToLower(h.Name)
		if name == "authorization" || name == "cookie" {
			return true
		}
	}
	return false
}
