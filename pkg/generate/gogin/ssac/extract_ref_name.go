//ff:func feature=gen-gogin type=util control=sequence
//ff:what extractRefName — SchemaRef에서 직접 $ref 또는 array items $ref 이름 추출

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractRefName returns the schema name referenced by propRef — either
// directly via $ref, or indirectly as the items type of an array schema.
// Returns "" when no ref is present.
func extractRefName(propRef *openapi3.SchemaRef) string {
	if propRef == nil {
		return ""
	}
	if propRef.Ref != "" {
		return refName(propRef.Ref)
	}
	if propRef.Value == nil {
		return ""
	}
	if !propRef.Value.Type.Is("array") {
		return ""
	}
	if propRef.Value.Items == nil || propRef.Value.Items.Ref == "" {
		return ""
	}
	return refName(propRef.Value.Items.Ref)
}
