//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what extractPathParams 가 path 템플릿 이름을 무변환으로 추출하는지 검증 (BUG-109)

package react

import (
	"reflect"
	"testing"
)

func TestExtractPathParams_PreservesOriginalName(t *testing.T) {
	cases := []struct {
		path string
		want []string
	}{
		{"/contracts/{contractId}/calculate-increase", []string{"contractId"}},
		{"/photos/{photoUrl}", []string{"photoUrl"}},
		{"/t/{tid}", []string{"tid"}},
		{"/api/items/{id}/sub/{subId}", []string{"id", "subId"}},
		{"/api/items", nil},
	}
	for _, c := range cases {
		got := extractPathParams(c.path)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("extractPathParams(%q) = %v, want %v", c.path, got, c.want)
		}
	}
}
