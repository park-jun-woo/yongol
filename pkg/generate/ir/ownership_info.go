//ff:type feature=gen-ir type=model
//ff:what OwnershipInfo -- Rego @ownership 어노테이션에서 추출한 소유권 조회 메타데이터

package ir

// OwnershipInfo carries the metadata needed by renderers to emit an ownership
// lookup query (e.g. OwnerLookup<Resource>) without re-interpreting the Rego
// policy. Extracted from rego.OwnershipMapping in convert_auth.go.
type OwnershipInfo struct {
	// Table is the DDL table name used for the ownership lookup query.
	Table string

	// OwnerColumn is the column name that identifies the owner (e.g. "owner_id").
	OwnerColumn string

	// ResourcePK is the primary key column of the resource table (e.g. "id").
	ResourcePK string
}
