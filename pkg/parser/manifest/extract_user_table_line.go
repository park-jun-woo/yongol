//ff:func feature=projectconfig type=parser control=iteration dimension=1
//ff:what extractUserTableLine — manifest.yaml 의 backend.auth.user_table 키 줄 번호 반환 (없으면 0)

package manifest

import (
	"gopkg.in/yaml.v3"
)

// extractUserTableLine walks the raw manifest.yaml bytes to backend.auth and
// returns the 1-based line number of the `user_table:` key. Returns 0 when
// the key is absent or the auth block cannot be located. Used by XDN-01 /
// XDN-02 to point diagnostics at the manifest line that needs editing.
func extractUserTableLine(data []byte) int {
	auth := FindAuthNode(data)
	if auth == nil || auth.Kind != yaml.MappingNode {
		return 0
	}
	for i := 0; i+1 < len(auth.Content); i += 2 {
		if auth.Content[i].Value == "user_table" {
			return auth.Content[i].Line
		}
	}
	return 0
}
