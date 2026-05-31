//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseArg / parseArgs / filterSubscribe / buildSubscribeInfo 단위 검증
package ssac

import (
	"testing"
)

func TestBuildSubscribeInfo(t *testing.T) {
	t.Run("with param", func(t *testing.T) {
		si := buildSubscribeInfo("order.completed", &ParamInfo{TypeName: "OnOrderCompletedMessage"})
		if si.Topic != "order.completed" || si.MessageType != "OnOrderCompletedMessage" {
			t.Errorf("si = %+v", si)
		}
	})
	t.Run("nil param", func(t *testing.T) {
		si := buildSubscribeInfo("x.y", nil)
		if si.Topic != "x.y" || si.MessageType != "" {
			t.Errorf("si = %+v", si)
		}
	})
}
