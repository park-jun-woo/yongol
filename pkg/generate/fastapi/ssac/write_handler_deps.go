//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeHandlerDeps — handler 함수의 의존성 주입 파라미터(current_user/session/event_bus) 출력

package ssac

import "strings"

// writeHandlerDeps writes the FastAPI dependency-injection parameters.
// current_user is skipped for pre-auth endpoints (login with @verify-password);
// event_bus is injected only when the plan contains @publish ops.
func writeHandlerDeps(b *strings.Builder, isPreAuth, needsEventBus bool) {
	if !isPreAuth {
		b.WriteString("    current_user: dict = Depends(get_current_user),\n")
	}
	b.WriteString("    session: AsyncSession = Depends(get_session),\n")
	if needsEventBus {
		b.WriteString("    event_bus: EventBus = Depends(get_event_bus),\n")
	}
}
