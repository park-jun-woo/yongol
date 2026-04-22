//ff:type feature=gen-gogin type=model
//ff:what ClaimField — manifest claims 에서 파싱된 단일 claim 필드

package auth

// ClaimField represents one JWT claim field derived from manifest.backend.auth.claims.
// Name is the Go struct field (PascalCase), Key is the JWT claim key, GoType is the Go type.
type ClaimField struct {
	Name   string // "ID", "Email", "Role", "OrgID"
	Key    string // "user_id", "email", "role", "org_id"
	GoType string // "int64", "string", "bool"
}
