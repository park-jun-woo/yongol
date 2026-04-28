//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what responseSchemaForStatus — status 에 해당하는 JSON response schema 선택

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// responseSchemaForStatus picks the JSON response schema associated
// with statusCode on route. Resolution order:
//  1. Exact status match (`200`, `201`, etc.).
//  2. First 2xx declared when statusCode is empty or its response has
//     no JSON content.
//  3. `default` response as a last resort.
//
// Returns nil when no declared response has a JSON schema. Callers use
// nil to mean "assertion cannot be evaluated; skip" rather than emitting
// a false positive.
func responseSchemaForStatus(route *apiRoute, statusCode string) *openapi3.Schema {
	if route == nil || route.Op == nil || route.Op.Responses == nil {
		return nil
	}
	resps := route.Op.Responses.Map()
	if statusCode != "" {
		if schema := jsonSchemaFromResponse(resps[statusCode]); schema != nil {
			return schema
		}
	}
	for code, r := range resps {
		if !strings.HasPrefix(code, "2") {
			continue
		}
		if schema := jsonSchemaFromResponse(r); schema != nil {
			return schema
		}
	}
	return jsonSchemaFromResponse(resps["default"])
}
