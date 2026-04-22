//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what indexFirst2xxResponse — responses 중 첫 번째 2xx 응답의 schema.properties 줄 번호를 색인
package openapi

import "gopkg.in/yaml.v3"

// indexFirst2xxResponse scans responses (MappingNode) for the first 2xx status
// code and records its schema.properties lines into ResponseFields.
func indexFirst2xxResponse(resps *yaml.Node, opID string, idx *LineIndex) {
	for k := 0; k+1 < len(resps.Content); k += 2 {
		code := resps.Content[k].Value
		if len(code) == 0 || code[0] != '2' {
			continue
		}
		props := schemaPropsOfBody(resps.Content[k+1])
		if props != nil {
			idx.ResponseFields[opID] = collectPropertyLines(props)
		}
		return
	}
}
