//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSortedKeys(t *testing.T) {
	m := openapi3.Schemas{"c": strSchema(), "a": strSchema(), "b": strSchema()}
	got := sortedKeys(m)
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedKeys = %v, want %v", got, want)
	}
}
