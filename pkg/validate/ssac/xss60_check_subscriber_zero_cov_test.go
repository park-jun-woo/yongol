//ff:func feature=validate type=test control=sequence
//ff:what TestXss60ByName_ZeroCov — XSS-60 publish/subscribe 타입 수집·비교 헬퍼 직접 호출
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60CheckSubscriber_ZeroCov(t *testing.T) {
	// no subscribe → nil.
	if d := xss60CheckSubscriber(parsessac.ServiceFunc{}, nil); d != nil {
		t.Errorf("non-subscriber should yield nil")
	}
	// subscriber with topic not in publish map → nil.
	fn := parsessac.ServiceFunc{
		Name:      "OnDone",
		Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
		Param:     &parsessac.ParamInfo{TypeName: "Msg"},
		Structs: []parsessac.StructInfo{{
			Name:   "Msg",
			Fields: []parsessac.StructField{{Name: "OrderID", Type: "int64"}},
		}},
	}
	if d := xss60CheckSubscriber(fn, map[string]map[string]string{}); d != nil {
		t.Errorf("topic absent from publishers should yield nil, got %v", d)
	}
	// matching topic, compatible field → no diag.
	pub := map[string]map[string]string{"order.completed": {"OrderID": "int64"}}
	if d := xss60CheckSubscriber(fn, pub); d != nil {
		t.Errorf("compatible field should yield no diag, got %v", d)
	}
}
