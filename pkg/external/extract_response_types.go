//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 메서드 목록에서 응답 타입 구조체를 추출한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func extractResponseTypes(serviceName string, methods []methodInfo, doc *openapi3.T) []structType {
	var types []structType

	for _, m := range methods {
		if m.ReturnType == "" {
			continue
		}

		op := findOperation(doc, m.HTTPMethod, m.Path)
		if op == nil {
			continue
		}

		resp, ok := op.Responses.Map()["200"]
		if !ok || resp.Value == nil {
			continue
		}
		ct := resp.Value.Content.Get("application/json")
		if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
			continue
		}

		st := structType{Name: m.ReturnType}
		propNames := sortedKeys(ct.Schema.Value.Properties)
		for _, name := range propNames {
			ref := ct.Schema.Value.Properties[name]
			st.Fields = append(st.Fields, structField{
				Name:     toPascalCase(name),
				GoType:   schemaToGoType(ref),
				JSONName: name,
			})
		}
		if len(st.Fields) > 0 {
			types = append(types, st)
		}
	}

	return types
}
