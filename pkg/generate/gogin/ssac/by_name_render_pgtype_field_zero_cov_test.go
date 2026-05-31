//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"strings"
	"testing"
)

func TestByNameRenderPgtypeField_ZeroCov(t *testing.T) {
	g := &methodGen{RespFields: map[string]responseField{
		"name": {JSONName: "name", IsRequired: true},
	}}
	// required → direct.
	if got := g.renderPgtypeField("name", "row.Name", "conv"); !strings.Contains(got, "Name: conv,") {
		t.Errorf("required pgtype field = %q", got)
	}
	// optional + not-already-pointer → ptrOf wrap.
	if got := g.renderPgtypeField("note", "row.Note", "conv"); !strings.Contains(got, "ptrOf(conv)") {
		t.Errorf("optional pgtype field = %q", got)
	}
}
