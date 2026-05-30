//ff:func feature=gen-gogin type=test control=branch topic=depth-report
//ff:what TestMeasureDepth_ParseError — 잘못된 소스에서 파서 에러 전파 검증

package qcheck

import "testing"

func TestMeasureDepth_ParseError(t *testing.T) {
	_, err := MeasureDepth("bad.go", "this is not valid go @@@")
	if err == nil {
		t.Fatalf("expected parse error for malformed source, got nil")
	}
}
