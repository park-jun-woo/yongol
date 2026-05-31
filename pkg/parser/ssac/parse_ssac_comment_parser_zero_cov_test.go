//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseBatch_ZeroCov — ssac 파서 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac

import (
	"testing"
)

func TestParseSsacCommentParser_ZeroCov(t *testing.T) {
	cp := &commentParser{}
	if err := cp.processLine("@get Course course = Course.FindByID({ID: request.id})", 1); err != nil {
		t.Errorf("processLine: %v", err)
	}
	cp.inResponse = true
	cp.processResponseBody("course: course")
}
