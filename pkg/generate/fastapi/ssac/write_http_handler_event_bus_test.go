//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteHTTPHandlerEventBus — 서브테스트 디스패치
package ssac

import "testing"

func TestWriteHTTPHandlerEventBus(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"PublishOpInjectsEventBus", subtestTestWriteHTTPHandlerEventBusPublishOpInjectsEventBus},
		{"NoPublishOpSkipsEventBus", subtestTestWriteHTTPHandlerEventBusNoPublishOpSkipsEventBus},
		{"EventBusAfterCurrentUser", subtestTestWriteHTTPHandlerEventBusEventBusAfterCurrentUser},
	} {
		t.Run(st.name, st.fn)
	}
}
