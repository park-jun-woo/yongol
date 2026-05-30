//ff:func feature=gen-gogin type=test control=branch topic=loop-report
//ff:what TestMeasurePureLines_ParseError — 잘못된 소스에서 파서 에러 전파 검증

package qcheck

import "testing"

func TestMeasurePureLines_ParseError(t *testing.T) {
	_, err := MeasurePureLines("bad.go", "@@@ not go")
	if err == nil {
		t.Fatalf("expected parse error for malformed source, got nil")
	}
}
