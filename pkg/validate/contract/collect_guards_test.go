//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestCollectGuards — 함수 body 내 len(x)/range x 언급을 guard map 에 기록 검증

package contract

import "testing"

func TestCollectGuards(t *testing.T) {
	body := mustBlock(t,
		"if len(xs) > 0 { _ = xs[0] }\n"+
			"for _, v := range ys { _ = v }\n"+
			"_ = zs[0]\n")
	guarded := map[string]bool{}
	collectGuards(body, guarded)

	if !guarded["xs"] {
		t.Fatalf("len(xs) should guard xs, got %v", guarded)
	}
	if !guarded["ys"] {
		t.Fatalf("range ys should guard ys, got %v", guarded)
	}
	if guarded["zs"] {
		t.Fatalf("zs is only indexed, not guarded; got %v", guarded)
	}
}
