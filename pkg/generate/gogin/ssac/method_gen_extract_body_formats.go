//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.extractBodyFormats — request body 스키마에서 format별 wrapper cast 정보 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractBodyFormats populates BodyFormats from the request body schema so
// mapValue can emit an explicit primitive cast (e.g. string(request.Body.Email))
// when the OpenAPI format produces an oapi-codegen wrapper type.
func (g *methodGen) extractBodyFormats(op *openapi3.Operation) {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return
	}
	mt := op.RequestBody.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return
	}
	// Populate BodyRequiredFields so wrapInsertExpr can distinguish
	// value types (required) from pointer types (optional) in oapi-codegen
	// output when choosing between Ptr / non-Ptr pgtypex variants (BUG-072).
	g.BodyRequiredFields = requiredSet(mt.Schema.Value)

	for propName, propRef := range mt.Schema.Value.Properties {
		if propRef == nil || propRef.Value == nil {
			continue
		}
		switch {
		case propRef.Value.Format != "":
			g.BodyFormats[propName] = propRef.Value.Format
		case propRef.Value.Type != nil && propRef.Value.Type.Is("string") && len(propRef.Value.Enum) > 0:
			// BUG-020 — oapi-codegen emits a named string type per
			// `enum:` string property (e.g. api.SignupJSONBodyPlanType).
			// Flag it so mapRequestValue wraps the access with a
			// `string(...)` cast when passed to sqlc params.
			g.BodyFormats[propName] = "enum"
		}
		// Record JSONB-shaped properties so sqlcArgs can hoist a
		// json.Marshal call for them (BUG-005 request direction).
		if isJSONBProperty(propRef) {
			g.BodyJSONBFields[propName] = true
		}
	}
}
