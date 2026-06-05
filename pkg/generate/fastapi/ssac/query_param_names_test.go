//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestQueryParamNames — queryParamNames 메타 목록에서 이름만 추출 검증
package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestQueryParamNames(t *testing.T) {
	in := []ir.QueryParamMeta{{Name: "limit"}, {Name: "cursor"}}
	got := queryParamNames(in)
	if !reflect.DeepEqual(got, []string{"limit", "cursor"}) {
		t.Errorf("queryParamNames() = %v, want [limit cursor]", got)
	}

	if empty := queryParamNames(nil); len(empty) != 0 {
		t.Errorf("queryParamNames(nil) = %v, want empty", empty)
	}
}
