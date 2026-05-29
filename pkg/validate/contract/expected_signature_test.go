//ff:func feature=validate-contract type=test control=sequence
//ff:what TestExpectedSignature — operationId 기준 기대 FuncSignature 계산 검증

package contract

import "testing"

func TestExpectedSignature(t *testing.T) {
	fs := buildFSWithOp("ActivateWorkflow")
	g := fs.Ground()

	t.Run("known opID", func(t *testing.T) {
		sig, ok := expectedSignature(g, "ActivateWorkflow")
		if !ok {
			t.Fatal("expected ok=true for known opID")
		}
		if sig.Name != "ActivateWorkflow" {
			t.Errorf("name = %q", sig.Name)
		}
		if len(sig.Params) != 2 || sig.Params[0].Type != "context.Context" {
			t.Errorf("params = %+v", sig.Params)
		}
		if sig.Params[1].Type != "api.ActivateWorkflowRequestObject" {
			t.Errorf("request param type = %q", sig.Params[1].Type)
		}
		if len(sig.Returns) != 2 || sig.Returns[1] != "error" || !sig.HasErr {
			t.Errorf("returns = %v hasErr=%v", sig.Returns, sig.HasErr)
		}
	})

	t.Run("unknown opID", func(t *testing.T) {
		if _, ok := expectedSignature(g, "Nope"); ok {
			t.Error("expected ok=false for unknown opID")
		}
	})

	t.Run("nil ground / empty opID", func(t *testing.T) {
		if _, ok := expectedSignature(nil, "X"); ok {
			t.Error("nil ground should be ok=false")
		}
		if _, ok := expectedSignature(g, ""); ok {
			t.Error("empty opID should be ok=false")
		}
	})
}
