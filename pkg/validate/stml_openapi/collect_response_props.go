//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectResponseProps — 응답 content 에서 property 이름+타입 수집

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

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
