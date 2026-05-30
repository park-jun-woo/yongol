//ff:func feature=validate type=test-helper control=sequence topic=openapi-structural
//ff:what o06EmptyLines — 빈 LineIndex 를 만든다(라인 폴백 0)

package openapi

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// o06EmptyLines returns a LineIndex with empty maps, so O-6 line resolution
// falls back to 0 for inline schemas in unit tests.
func o06EmptyLines() *oapiparser.LineIndex {
	return &oapiparser.LineIndex{
		Paths:      map[string]int{},
		Operations: map[string]int{},
		Schemas:    map[string]int{},
	}
}
