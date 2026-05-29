//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractResponseArrayItemFields_NoPaths — paths 없는 doc 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseArrayItemFields_NoPaths(t *testing.T) {
	doc := &openapi3.T{}
	result := ExtractResponseArrayItemFields(doc)
	if len(result) != 0 {
		t.Errorf("expected empty map for doc without paths, got %v", result)
	}
}
