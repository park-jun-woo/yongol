//ff:func feature=validate type=util control=sequence dimension=1 topic=response-body-required
//ff:what opIDOrPath — operationId 우선, 없으면 path 템플릿 fallback

package openapi

// opIDOrPath returns operationId when present, otherwise falls back to the
// path template so an O-5 diagnostic still locates the offending operation
// even when O-4 (operationId required) has not yet been resolved.
func opIDOrPath(opID, path string) string {
	if opID != "" {
		return opID
	}
	return path
}
