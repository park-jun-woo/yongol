//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseArg / parseArgs / filterSubscribe / buildSubscribeInfo 단위 검증
package ssac

import (
	"testing"
)

func TestFilterSubscribe(t *testing.T) {
	sf := &ServiceFunc{Param: &ParamInfo{TypeName: "Msg"}}
	seqs := []Sequence{
		{Type: "subscribe", Topic: "order.done"},
		{Type: "get", Model: "Course.FindByID"},
	}
	filtered := filterSubscribe(sf, seqs)
	if len(filtered) != 1 || filtered[0].Type != "get" {
		t.Errorf("filtered = %+v", filtered)
	}
	if sf.Subscribe == nil || sf.Subscribe.Topic != "order.done" || sf.Subscribe.MessageType != "Msg" {
		t.Errorf("subscribe = %+v", sf.Subscribe)
	}
}
