//ff:func feature=agent type=test control=sequence
//ff:what TestDiagMessages — 진단 목록에서 메시지 문자열을 순서대로 추출 검증
package agent

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagMessages(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Message: "first"},
		{Message: "second"},
	}
	got := diagMessages(diags)
	want := []string{"first", "second"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("diagMessages = %v, want %v", got, want)
	}
	if got := diagMessages(nil); len(got) != 0 {
		t.Errorf("diagMessages(nil) = %v, want empty", got)
	}
}
