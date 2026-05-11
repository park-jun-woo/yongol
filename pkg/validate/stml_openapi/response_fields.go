//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what responseFields — Operation의 200 응답 스키마에서 top-level property → type 맵 추출

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// responseFields extracts top-level property names and their types from
// the first 2xx response schema of the given operation. allOf members
// are resolved one level deep.
func responseFields(op *openapi3.Operation) map[string]responseFieldInfo {
	out := make(map[string]responseFieldInfo)
	if op == nil || op.Responses == nil {
		return out
	}
	// Try explicit 2xx status codes first.
	for _, code := range []string{"200", "201"} {
		resp := op.Responses.Status(statusInt(code))
		if resp != nil && resp.Value != nil {
			collectResponseProps(out, resp.Value)
			if len(out) > 0 {
				return out
			}
		}
	}
	return out
}

// statusInt converts a status code string to int. Only used for 2xx.
func statusInt(code string) int {
	switch code {
	case "200":
		return 200
	case "201":
		return 201
	default:
		return 0
	}
}

// collectResponseProps fills out with property names from the response content schema.
func collectResponseProps(out map[string]responseFieldInfo, resp *openapi3.Response) {
	if resp.Content == nil {
		return
	}
	for _, mt := range resp.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		addSchemaProps(out, mt.Schema.Value)
		return
	}
}

// addSchemaProps adds top-level properties and resolves allOf one level.
func addSchemaProps(out map[string]responseFieldInfo, s *openapi3.Schema) {
	if s == nil {
		return
	}
	for name, ref := range s.Properties {
		out[name] = responseFieldInfo{typ: schemaType(ref)}
	}
	for _, allOfRef := range s.AllOf {
		if allOfRef == nil || allOfRef.Value == nil {
			continue
		}
		for name, ref := range allOfRef.Value.Properties {
			out[name] = responseFieldInfo{typ: schemaType(ref)}
		}
	}
}

// schemaType returns the primary type string from a schema ref, or "" if unavailable.
func schemaType(ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil || ref.Value.Type == nil {
		return ""
	}
	types := ref.Value.Type.Slice()
	if len(types) == 0 {
		return ""
	}
	return types[0]
}
