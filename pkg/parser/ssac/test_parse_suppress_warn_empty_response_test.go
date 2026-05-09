//ff:func feature=ssac-parse type=test control=sequence topic=response
//ff:what @response! 빈 응답 파싱 검증 — nil 아닌 빈 Sequence 반환 (BUG-053)

package ssac

import "testing"

// TestParseSuppressWarnEmptyResponse verifies that @response! with no args
// returns a non-nil Sequence (Type=response, SuppressWarn=true, Target=""),
// so that buildResponse is invoked and a return statement is generated.
func TestParseSuppressWarnEmptyResponse(t *testing.T) {
	src := `package service

// @delete! Webhook.Delete({ID: request.ID})
// @response!
func DeleteWebhook() {}
`
	sfs := parseTestFile(t, src)
	var resp *Sequence
	for i := range sfs[0].Sequences {
		if sfs[0].Sequences[i].Type == SeqResponse {
			resp = &sfs[0].Sequences[i]
			break
		}
	}
	if resp == nil {
		t.Fatal("expected response sequence for @response! (was nil before fix)")
	}
	if !resp.SuppressWarn {
		t.Error("expected SuppressWarn=true")
	}
	if resp.Target != "" {
		t.Errorf("expected empty Target, got %q", resp.Target)
	}
	if len(resp.Fields) != 0 {
		t.Errorf("expected empty Fields, got %v", resp.Fields)
	}
}
