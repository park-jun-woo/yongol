//ff:type feature=validate type=model topic=openapi-structural
//ff:what o06SchemaEntry — O-6 walk 가 수집하는 (스키마 값, components 스키마명) 쌍

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// o06SchemaEntry pairs a resolved schema value with the components schema name
// it belongs to (empty for inline request/response schemas). schemaName drives
// LineIndex lookups; it is empty for inline schemas because LineIndex has no
// per-inline-schema index, in which case o06CheckSchemaRequired falls back to
// line 0.
type o06SchemaEntry struct {
	schema     *openapi3.Schema
	schemaName string
}
