//ff:func feature=projectconfig type=parser control=sequence
//ff:what extractAuthLines — manifest.yaml 원본을 순회해 backend.auth.claims/roles 의 줄 번호를 수집
package manifest

// extractAuthLines parses the raw yaml into a generic yaml.Node tree and walks
// to backend.auth.{claims, roles}, returning per-key/per-value 1-based line
// numbers. Missing entries map to 0.
func extractAuthLines(data []byte) (claimLines, roleLines map[string]int) {
	claimLines = map[string]int{}
	roleLines = map[string]int{}

	auth := FindAuthNode(data)
	if auth == nil {
		return
	}
	collectClaimLines(auth, claimLines)
	collectRoleLines(auth, roleLines)
	return
}
