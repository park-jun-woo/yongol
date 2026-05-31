//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"strings"
	"testing"
)

func TestByNameBuildResponseConvert_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "GetWorkflow", SuccessStatus: 200}
	// scalar model.
	lines := g.buildResponseConvert("Workflow", "wf")
	if len(lines) != 3 || !strings.Contains(lines[0], "convertWorkflow(wf)") {
		t.Errorf("scalar convert = %v", lines)
	}
	// list model.
	listLines := g.buildResponseConvert("[]Workflow", "wfs")
	if !strings.Contains(listLines[0], "convertWorkflowList(wfs)") {
		t.Errorf("list convert = %v", listLines)
	}
}
