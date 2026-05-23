//ff:func feature=agent type=adapter control=sequence
//ff:what verifyOpenAPI — kin-openapi loader로 YAML 구조 검증

package agent

import "github.com/getkin/kin-openapi/openapi3"

func verifyOpenAPI(yamlData []byte) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(yamlData)
	if err != nil {
		return err
	}
	if err := doc.Validate(loader.Context); err != nil {
		return err
	}
	return nil
}
