//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what schemaPropertyNames — schema top-level + allOf members 의 property 이름 집합

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// schemaPropertyNames collects top-level property keys from an OpenAPI
// schema, resolving one level of allOf so discriminated schemas work.
func schemaPropertyNames(s *openapi3.Schema) map[string]struct{} {
	out := map[string]struct{}{}
	if s == nil {
		return out
	}
	for name := range s.Properties {
		out[name] = struct{}{}
	}
	addAllOfPropertyNames(out, s.AllOf)
	return out
}
