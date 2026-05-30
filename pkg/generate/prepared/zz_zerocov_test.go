//ff:func feature=generate type=test
//ff:what zz_zerocov_test — prepared.funcUsesQueue 0% 커버리지 단위 테스트
package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestFuncUsesQueue_ZeroCov(t *testing.T) {
	// Subscribe set → true (first branch).
	sub := ssac.ServiceFunc{Subscribe: &ssac.SubscribeInfo{Topic: "t"}}
	if !funcUsesQueue(sub) {
		t.Error("expected true for Subscribe set")
	}

	// publish sequence → true (loop branch).
	pub := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "publish"}}}
	if !funcUsesQueue(pub) {
		t.Error("expected true for publish sequence")
	}

	// No subscribe, no publish → false.
	none := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "response"}}}
	if funcUsesQueue(none) {
		t.Error("expected false for non-queue func")
	}

	// Empty func → false.
	if funcUsesQueue(ssac.ServiceFunc{}) {
		t.Error("expected false for empty func")
	}
}
