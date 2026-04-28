//ff:func feature=validate type=util control=iteration dimension=1 topic=policy-check
//ff:what ownerLookupName — @ownership resource (snake_case) → OwnerLookup<Pascal> 쿼리명

package query_rego

import "strings"

// ownerLookupName builds the canonical sqlc query name for the supplied
// `@ownership` resource. Resource strings may arrive as snake_case
// (execution_log) or lowercase (workflow); either way the PascalCase form
// is produced by capitalising each segment.
func ownerLookupName(resource string) string {
	parts := strings.Split(resource, "_")
	var b strings.Builder
	b.WriteString("OwnerLookup")
	for _, part := range parts {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			b.WriteString(part[1:])
		}
	}
	return b.String()
}
