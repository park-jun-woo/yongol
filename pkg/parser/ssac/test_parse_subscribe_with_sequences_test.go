//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @subscribe 복수 시퀀스 파싱 검증 — get/call/put 3단계

package ssac

import "testing"

func TestParseSubscribeWithSequences(t *testing.T) {
	src := `package service

type OnOrderCompletedMessage struct {
	OrderID int64
	Email   string
}

// @subscribe "order.completed"
// @get Order order = Order.FindByID({ID: message.OrderID})
// @call mail.SendEmail({To: message.Email, Subject: "done"})
// @put Order.UpdateNotified({ID: order.ID, Notified: "true"})
func OnOrderCompleted(message OnOrderCompletedMessage) {}
`
	sfs := parseTestFile(t, src)
	sf := sfs[0]
	if sf.Subscribe == nil {
		t.Fatal("expected Subscribe to be set")
	}
	assertEqual(t, "Subscribe.Topic", sf.Subscribe.Topic, "order.completed")
	assertEqual(t, "Subscribe.MessageType", sf.Subscribe.MessageType, "OnOrderCompletedMessage")
	if len(sf.Sequences) != 3 {
		t.Fatalf("expected 3 sequences, got %d", len(sf.Sequences))
	}
	assertEqual(t, "seq0.Type", sf.Sequences[0].Type, SeqGet)
	assertEqual(t, "seq1.Type", sf.Sequences[1].Type, SeqCall)
	assertEqual(t, "seq2.Type", sf.Sequences[2].Type, SeqPut)
}
