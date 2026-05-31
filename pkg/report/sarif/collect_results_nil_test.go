//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestCollectResults — nil 리포트 / 다중 step / non-error,warning 필터 / fired 집합 검증
package sarif

import (
	"testing"
)

func TestCollectResults_Nil(t *testing.T) {
	results, fired := collectResults(nil, "", nil)
	if results == nil || len(results) != 0 {
		t.Errorf("results: want empty non-nil, got %+v", results)
	}
	if fired == nil || len(fired) != 0 {
		t.Errorf("fired: want empty non-nil, got %+v", fired)
	}
}
