//ff:type feature=manifest type=model
//ff:what LineIndex — per-node line number index built by raw yaml.v3 parsing of an OpenAPI file
package openapi

// LineIndex maps OpenAPI nodes to their 1-based line numbers in the source
// YAML file. Because kin-openapi/openapi3 does not expose line information,
// the same file is parsed a second time with yaml.v3 to build this index.
//
// All line numbers reference the line where the node's *key* appears (1-based).
// 0 means unknown.
type LineIndex struct {
	File string

	// Operations: operationId → line of that operation block (the line where
	// the operationId key appears). Indexed by operationId, not by path/method.
	Operations map[string]int

	// OperationFields: operationId → field name → line.
	// Line where each property under
	// requestBody.content.application/json.schema.properties is declared.
	RequestFields map[string]map[string]int

	// ResponseFields: operationId → field name → line.
	// First 2xx response.content.application/json.schema.properties.
	ResponseFields map[string]map[string]int

	// Schemas: schema name → line of the schema block under components.schemas.
	Schemas map[string]int

	// SchemaProperties: schema name → property name → line.
	SchemaProperties map[string]map[string]int

	// Paths: path template → line of the path key (e.g. "/users/{id}").
	Paths map[string]int
}
