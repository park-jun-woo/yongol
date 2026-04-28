//ff:func feature=gen-gogin type=util control=sequence
//ff:what ownerLookupQueryName — 공용 sqlc 쿼리 이름 "OwnerLookup<PascalResource>" 생성

package ssac

// ownerLookupQueryName builds the canonical sqlc query name
// (`OwnerLookup<PascalResource>`) shared with XQP-30 validate. Keep in
// lockstep with pkg/validate/query_rego/xqp_30_owner_lookup_query.go.
func ownerLookupQueryName(resource string) string {
	return "OwnerLookup" + pascalCase(resource)
}
