//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @call ErrStatus + Result 동시 파싱 검증

package ssac

import "testing"

func TestParseCallErrStatusWithResult(t *testing.T) {
	src := `package service

// @call ChargeResult charge = billing.Charge({Amount: order.Total}) 402
func PlaceOrder(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqCall)
	assertEqual(t, "Model", seq.Model, "billing.Charge")
	if seq.Result == nil {
		t.Fatal("expected result")
	}
	assertEqual(t, "Result.Type", seq.Result.Type, "ChargeResult")
	if seq.ErrStatus != 402 {
		t.Errorf("expected ErrStatus 402, got %d", seq.ErrStatus)
	}
}
