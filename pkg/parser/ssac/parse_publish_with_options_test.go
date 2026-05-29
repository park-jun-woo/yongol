//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @publish Options 파싱 검증 — {delay: 1800} 옵션 확인

package ssac

import "testing"

func TestParsePublishWithOptions(t *testing.T) {
	src := `package service

// @publish "cart.abandoned" {CartID: cart.ID, UserID: currentUser.ID} {delay: 1800}
func AbandonCart() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqPublish)
	assertEqual(t, "Topic", seq.Topic, "cart.abandoned")
	assertEqual(t, "Inputs.CartID", seq.Inputs["CartID"], "cart.ID")
	if seq.Options == nil {
		t.Fatal("expected options")
	}
	assertEqual(t, "Options.delay", seq.Options["delay"], "1800")
}
