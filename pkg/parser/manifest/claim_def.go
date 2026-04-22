//ff:type feature=projectconfig type=model
//ff:what JWT 클레임 정의 구조체
package manifest

// ClaimDef describes a single JWT claim with its key and Go type.
type ClaimDef struct {
	Key        string // JWT claim key (e.g. "org_id")
	GoType     string // Go type (e.g. "int64"), default "string"
	SourceLine int    // 1-based line number of the claim entry in manifest.yaml (0 = unknown)
}
