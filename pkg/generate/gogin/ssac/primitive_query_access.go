//ff:func feature=gen-gogin type=util control=sequence
//ff:what primitiveQueryAccess — required면 accessor 그대로, 아니면 deref* 래핑

package ssac

// primitiveQueryAccess returns accessor unchanged when the query parameter
// is required (oapi-codegen emits a non-pointer field for required primitive
// queries) and wraps it in the matching deref* helper otherwise.
func primitiveQueryAccess(required bool, accessor, derefFn string) string {
	if required {
		return accessor
	}
	return derefFn + "(" + accessor + ")"
}
