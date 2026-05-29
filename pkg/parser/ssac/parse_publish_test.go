//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @publish 파싱 검증 — Topic, Inputs, Options nil 확인

package ssac

import "testing"

func TestParsePublish(t *testing.T) {
	src := `package service

// @get Order order = Order.FindByID({ID: request.OrderID})
// @publish "order.completed" {OrderID: order.ID, Email: order.Email}
// @response { order: order }
func CompleteOrder() {}
`
	sfs := parseTestFile(t, src)
	if len(sfs[0].Sequences) != 3 {
		t.Fatalf("expected 3 sequences, got %d", len(sfs[0].Sequences))
	}
	seq := sfs[0].Sequences[1]
	assertEqual(t, "Type", seq.Type, SeqPublish)
	assertEqual(t, "Topic", seq.Topic, "order.completed")
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs.OrderID", seq.Inputs["OrderID"], "order.ID")
	assertEqual(t, "Inputs.Email", seq.Inputs["Email"], "order.Email")
	if seq.Options != nil {
		t.Errorf("expected nil options, got %v", seq.Options)
	}
}
