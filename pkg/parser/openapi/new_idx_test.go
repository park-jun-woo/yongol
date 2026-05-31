//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

func newIdx() *LineIndex {
	return &LineIndex{
		Operations:       map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		Schemas:          map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Paths:            map[string]int{},
	}
}
