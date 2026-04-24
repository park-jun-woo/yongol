//ff:func feature=validate type=util control=sequence topic=hurl-openapi
//ff:what responseSchemaForStatus + jsonPathReachable — XOH-04/08 공용 schema 탐색

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
		if strings.HasPrefix(code, "2") {
			if schema := jsonSchemaFromResponse(r); schema != nil {
				return schema
			}
		}
	}
	return jsonSchemaFromResponse(resps["default"])
}

// jsonSchemaFromResponse extracts the first JSON-ish media type's
// schema from a response ref, treating absent values as nil.
func jsonSchemaFromResponse(r *openapi3.ResponseRef) *openapi3.Schema {
	if r == nil || r.Value == nil {
		return nil
	}
	for ct, mt := range r.Value.Content {
		if !strings.Contains(strings.ToLower(ct), "json") {
			continue
		}
		if mt == nil || mt.Schema == nil {
			continue
		}
		return mt.Schema.Value
	}
	return nil
}

// jsonPathReachable walks a dotted JSONPath (`$.user.id`,
// `$.items[0].name`) through an OpenAPI schema and returns true when
// every segment resolves. Array indexes (`[n]`) descend into the
// `items` schema; unknown segments short-circuit to false.
//
// Wildcards (`$..email`, `$[*]`) are conservatively treated as
// reachable — hurl users often lean on them for flexible assertions
// and a strict walker would produce noisy false positives.
func jsonPathReachable(path string, schema *openapi3.Schema) bool {
	if schema == nil || path == "" {
		return false
	}
	if strings.Contains(path, "..") || strings.Contains(path, "[*]") || strings.Contains(path, "[?") {
		return true
	}
	segs := parseJSONPath(path)
	cur := schema
	for _, seg := range segs {
		cur = descend(cur, seg)
		if cur == nil {
			return false
		}
	}
	return true
}

// parseJSONPath splits `$.a.b[0].c` into `["a", "b", "[0]", "c"]`.
func parseJSONPath(path string) []string {
	p := strings.TrimPrefix(path, "$")
	p = strings.TrimPrefix(p, ".")
	var out []string
	var cur strings.Builder
	for i := 0; i < len(p); i++ {
		c := p[i]
		switch c {
		case '.':
			if cur.Len() > 0 {
				out = append(out, cur.String())
				cur.Reset()
			}
		case '[':
			if cur.Len() > 0 {
				out = append(out, cur.String())
				cur.Reset()
			}
			end := strings.Index(p[i:], "]")
			if end < 0 {
				return out
			}
			out = append(out, p[i:i+end+1])
			i += end
		default:
			cur.WriteByte(c)
		}
	}
	if cur.Len() > 0 {
		out = append(out, cur.String())
	}
	return out
}

// descend advances one segment through a schema. Array segments pass
// through to `items`; object segments look up `properties`.
func descend(s *openapi3.Schema, seg string) *openapi3.Schema {
	if s == nil {
		return nil
	}
	if strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]") {
		if s.Items != nil && s.Items.Value != nil {
			return s.Items.Value
		}
		return nil
	}
	if prop, ok := s.Properties[seg]; ok {
		if prop != nil {
			return prop.Value
		}
	}
	for _, ref := range s.AllOf {
		if ref == nil || ref.Value == nil {
			continue
		}
		if prop, ok := ref.Value.Properties[seg]; ok && prop != nil {
			return prop.Value
		}
	}
	return nil
}
