//ff:func feature=external type=generator control=sequence
//ff:what OpenAPI 데이터를 바이트로 받아 코드를 생성하여 반환한다 (테스트용)
package external

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

// GenerateBytes generates code and returns it as bytes (for testing).
func GenerateBytes(source string, data []byte) ([]byte, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("parse OpenAPI: %w", err)
	}
	if err := doc.Validate(loader.Context); err != nil {
		return nil, fmt.Errorf("validate OpenAPI: %w", err)
	}
	serviceName := inferServiceName(source, doc)
	return generateCode(serviceName, doc)
}
