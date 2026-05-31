//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSortedPathKeys(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/zebra", &openapi3.PathItem{})
	doc.Paths.Set("/apple", &openapi3.PathItem{})
	doc.Paths.Set("/mango", &openapi3.PathItem{})
	got := sortedPathKeys(doc)
	want := []string{"/apple", "/mango", "/zebra"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedPathKeys = %v, want %v", got, want)
	}
}
