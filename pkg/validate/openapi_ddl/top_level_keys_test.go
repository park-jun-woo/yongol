//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what topLevelKeys — nil ref/value 는 nil, inline 객체는 정렬된 top-level 키 슬라이스 반환

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestTopLevelKeys(t *testing.T) {
	if got := topLevelKeys(nil); got != nil {
		t.Errorf("nil ref keys = %v, want nil", got)
	}
	if got := topLevelKeys(&openapi3.SchemaRef{}); got != nil {
		t.Errorf("nil value keys = %v, want nil", got)
	}
	got := topLevelKeys(inlineRef("name", "id", "created_at"))
	want := []string{"created_at", "id", "name"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("keys = %v, want %v (sorted)", got, want)
	}
}
