//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTryExtractFromPathItem_NoVerbMatch_ZeroCov — verb 미매칭 시 false
package ssac

import (
	"testing"
)

func TestTryExtractFromPathItem_NoVerbMatch_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	doc := docZeroCov("GetWidget")
	pathItem := doc.Paths.Find("/widgets/{id}")
	if g.tryExtractFromPathItem(pathItem, "OtherOp") {
		t.Errorf("expected false for non-matching operationId")
	}
}
