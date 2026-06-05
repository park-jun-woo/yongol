//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHandlerDeps — writeHandlerDeps current_user/session/event_bus 의존성 출력 분기 검증
package ssac

import (
	"strings"
	"testing"
)

func TestWriteHandlerDeps(t *testing.T) {
	t.Run("FullDeps", func(t *testing.T) {
		var b strings.Builder
		writeHandlerDeps(&b, false, true)
		got := b.String()
		for _, want := range []string{"current_user", "get_session", "event_bus"} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q in output, got %q", want, got)
			}
		}
	})

	t.Run("PreAuthNoEventBus", func(t *testing.T) {
		var b strings.Builder
		writeHandlerDeps(&b, true, false)
		got := b.String()
		if strings.Contains(got, "current_user") {
			t.Errorf("pre-auth should skip current_user, got %q", got)
		}
		if strings.Contains(got, "event_bus") {
			t.Errorf("no event_bus expected, got %q", got)
		}
		if !strings.Contains(got, "get_session") {
			t.Errorf("session always present, got %q", got)
		}
	})
}
