//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func byNameO06FS(doc *openapi3.T) *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPIDoc: doc,
		OpenAPILines: &oapiparser.LineIndex{
			Schemas:          map[string]int{"Workflow": 10},
			SchemaProperties: map[string]map[string]int{"Workflow": {"name": 11}},
			Paths:            map[string]int{},
			Operations:       map[string]int{},
		},
	}
}
