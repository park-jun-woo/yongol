//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-manifest
//ff:what isKnownRefPath — manifest 참조 경로가 지원 화이트리스트에 포함되는지 확인

package ssac_manifest

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// isKnownRefPath checks whether a manifest ref path is in the supported whitelist.
func isKnownRefPath(path string) bool {
	for _, known := range manifest.KnownRefPaths() {
		if known == path {
			return true
		}
	}
	return false
}
