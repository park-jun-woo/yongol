//ff:func feature=generate type=test-helper control=sequence
//ff:what 테스트용 OpenAPI 문서를 로드하는 헬퍼 함수

package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func loadTestDoc(t *testing.T) *openapi3.T {
	t.Helper()
	yaml := []byte(`openapi: 3.0.3
info:
  title: t
  version: "0"
paths:
  /items:
    post:
      operationId: CreateItem
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [title]
              properties:
                title: { type: string }
                count: { type: integer }
`)
	doc, err := openapi3.NewLoader().LoadFromData(yaml)
	if err != nil {
		t.Fatalf("load test doc: %v", err)
	}
	return doc
}
