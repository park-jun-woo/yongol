//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"testing"
)

func TestByNameApplyOperation_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("GetWidget")
	doc := docZeroCov("GetWidget")
	pathItem := doc.Paths.Find("/widgets/{id}")
	v := verbOp{method: "GET", op: pathItem.Get}
	g.applyOperation(pathItem, v, "GetWidget")
	if g.Method != "GET" {
		t.Errorf("applyOperation Method = %q", g.Method)
	}
	if !g.PathParams["id"] {
		t.Errorf("applyOperation did not load path params: %v", g.PathParams)
	}
}
