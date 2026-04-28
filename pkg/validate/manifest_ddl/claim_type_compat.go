//ff:func feature=validate type=util control=selection topic=manifest-infra
//ff:what claimTypeCompatible — claim Go 타입 (int64/string/bool) 이 DDL 컬럼 Go 타입과 정합

package manifest_ddl

// claimTypeCompatible reports whether a manifest claim's Go type is
// compatible with the DDL column's Go type produced by the ddl parser
// (`pgTypeToGo`). The comparison is at the Go-type level because the parser
// already collapses BIGINT/INTEGER → int64, VARCHAR/TEXT/UUID/CHAR →
// string, BOOLEAN/BOOL → bool. XDN-04 piggy-backs on that mapping rather
// than re-parsing raw SQL.
//
// Empty claimGoType is treated as "string" (the ClaimDef default applied
// by parseRawClaims, but defended here for callers that bypass it).
func claimTypeCompatible(claimGoType, ddlGoType string) bool {
	switch claimGoType {
	case "":
		return "string" == ddlGoType
	default:
		return claimGoType == ddlGoType
	}
}
