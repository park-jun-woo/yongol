//ff:func feature=gen-react type=test control=sequence
//ff:what captureDeclaredOps — capture+opID 만 수집, opID 없거나 capture 없는 action 은 제외 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCaptureDeclaredOps(t *testing.T) {
	pages := []stml.PageSpec{
		{Actions: []stml.ActionBlock{
			// capture + operationId -> collected
			{OperationID: "Login", Captures: []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}},
			// captures present but operationId empty -> excluded
			{OperationID: "", Captures: []stml.CaptureBind{{RespField: "x", Sink: "auth.token"}}},
			// operationId present but no captures -> excluded
			{OperationID: "CreateThing"},
		}},
		{Actions: []stml.ActionBlock{
			{OperationID: "SignUp", Captures: []stml.CaptureBind{{RespField: "refresh_token", Sink: "auth.refresh"}}},
		}},
	}

	ops := captureDeclaredOps(pages)

	if len(ops) != 2 {
		t.Fatalf("collected = %v, want exactly Login and SignUp", ops)
	}
	if !ops["Login"] || !ops["SignUp"] {
		t.Errorf("missing capture-declared op: %v", ops)
	}
	if ops["CreateThing"] {
		t.Errorf("CreateThing has no captures, must not be collected")
	}
	if ops[""] {
		t.Errorf("empty operationId must not be collected")
	}

	// no pages -> empty (non-nil) set
	if got := captureDeclaredOps(nil); got == nil || len(got) != 0 {
		t.Errorf("nil pages = %v, want empty set", got)
	}
}
