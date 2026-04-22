//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what stripCRUDPrefix — operationId에서 CRUD 동사 prefix 제거

package openapi_ddl

import "strings"

// stripCRUDPrefix removes leading verbs commonly used in operationIds.
func stripCRUDPrefix(opID string) string {
	prefixes := []string{"list", "get", "create", "update", "delete", "patch", "post", "put", "fetch", "search", "find"}
	lower := strings.ToLower(opID)
	for _, p := range prefixes {
		if strings.HasPrefix(lower, p) && len(opID) > len(p) {
			rest := opID[len(p):]
			if rest != "" && rest[0] >= 'A' && rest[0] <= 'Z' {
				return rest
			}
		}
	}
	return opID
}
