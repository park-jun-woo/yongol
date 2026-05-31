//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseArg / parseArgs / filterSubscribe / buildSubscribeInfo 단위 검증
package ssac

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	got := parseArgs(`request.ID, "admin" , 7`)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3 (%+v)", len(got), got)
	}
	if got[0].Source != "request" || got[0].Field != "ID" {
		t.Errorf("arg0 = %+v", got[0])
	}
	if got[1].Literal != "admin" || !got[1].IsQuoted {
		t.Errorf("arg1 = %+v", got[1])
	}
	if got[2].Literal != "7" {
		t.Errorf("arg2 = %+v", got[2])
	}
	// empty parts skipped
	if g := parseArgs("  ,  "); g != nil {
		t.Errorf("empty args = %+v, want nil", g)
	}
}
