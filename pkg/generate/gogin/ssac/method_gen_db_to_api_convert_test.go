//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.dbToAPIConvert 단위 테스트 (convert<Ref>(<var>) 표현식 생성)
package ssac

import (
	"testing"
)

func TestMethodGenDbToAPIConvert(t *testing.T) {
	g := &methodGen{}
	if got := g.dbToAPIConvert("row", "Workflow"); got != "convertWorkflow(row)" {
		t.Errorf("dbToAPIConvert = %q, want %q", got, "convertWorkflow(row)")
	}
	if got := g.dbToAPIConvert("u", "User"); got != "convertUser(u)" {
		t.Errorf("dbToAPIConvert = %q, want %q", got, "convertUser(u)")
	}
}
