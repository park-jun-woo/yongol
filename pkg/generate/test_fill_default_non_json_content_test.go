//ff:func feature=generate type=test control=sequence
//ff:what application/json이 아닌 content type의 requestBody도 기본 타입으로 채우는지 검증

package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefault_NonJSONContentType(t *testing.T) {
	yaml := []byte(`openapi: 3.0.3
info:
  title: t
  version: "0"
paths:
  /upload:
    post:
      operationId: UploadFile
      requestBody:
        required: true
        content:
          multipart/form-data:
            schema:
              type: object
              required: [file]
              properties:
                file: { type: string, format: binary }
                description: { type: string }
      responses:
        '200':
          description: OK
`)
	doc, err := openapi3.NewLoader().LoadFromData(yaml)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	pages := []stmlparser.PageSpec{
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "UploadFile",
					Fields: []stmlparser.FieldBind{
						{Name: "file"},
						{Name: "description"},
					},
				},
			},
		},
	}
	result := fillDefaultRequestConstraints(pages, doc, nil)
	fields, ok := result["UploadFile"]
	if !ok {
		t.Fatal("UploadFile not found in result")
	}
	if got := fields["file"].Type; got != "string" {
		t.Errorf("file.Type = %q, want %q", got, "string")
	}
	if got := fields["file"].Format; got != "binary" {
		t.Errorf("file.Format = %q, want %q", got, "binary")
	}
	if got := fields["description"].Type; got != "string" {
		t.Errorf("description.Type = %q, want %q", got, "string")
	}
	if got := fields["file"].Required; !got {
		t.Error("file.Required = false, want true")
	}
}
