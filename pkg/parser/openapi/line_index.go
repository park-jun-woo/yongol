//ff:type feature=manifest type=model
//ff:what LineIndex — yaml.v3 raw parse 로 OpenAPI 노드별 줄 번호 색인
package openapi

// LineIndex maps OpenAPI nodes to their 1-based line numbers in the source
// YAML file. kin-openapi/openapi3 는 line 정보를 제공하지 않으므로 같은 파일을
// yaml.v3 로 한번 더 parsing 해서 별도로 구축한다.
//
// 모든 줄 번호는 해당 노드의 *key* 가 등장한 줄을 기준으로 한다 (1-based).
// 줄 번호 0 은 "미상" 을 의미한다.
type LineIndex struct {
	File string

	// Operations: operationId → line of that operation block (operationId
	// 키가 등장한 줄). path/method 가 아닌 operationId 단위로 색인.
	Operations map[string]int

	// OperationFields: operationId → field name → line.
	// requestBody.content.application/json.schema.properties 의 각 property
	// 가 선언된 줄.
	RequestFields map[string]map[string]int

	// ResponseFields: operationId → field name → line.
	// 첫 번째 2xx response.content.application/json.schema.properties.
	ResponseFields map[string]map[string]int

	// Schemas: schema name → line of the schema block under components.schemas.
	Schemas map[string]int

	// SchemaProperties: schema name → property name → line.
	SchemaProperties map[string]map[string]int

	// Paths: path template → line of the path key (e.g. "/users/{id}").
	Paths map[string]int
}
