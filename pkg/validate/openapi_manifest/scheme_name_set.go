//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what schemeNameSet — OpenAPI components.securitySchemes 이름 set 수집

package openapi_manifest

import "github.com/getkin/kin-openapi/openapi3"

func schemeNameSet(doc *openapi3.T) map[string]bool {
	set := map[string]bool{}
	if doc == nil || doc.Components == nil {
		return set
	}
	for name := range doc.Components.SecuritySchemes {
		set[name] = true
	}
	return set
}
