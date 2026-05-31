//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestByName_ZeroCov — openapi 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package openapi

func newLineIndexForTest() *LineIndex {
	return &LineIndex{
		Operations:       map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		Schemas:          map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Paths:            map[string]int{},
	}
}
