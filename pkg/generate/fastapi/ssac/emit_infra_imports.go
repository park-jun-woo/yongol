//ff:func feature=gen-fastapi type=util control=sequence
//ff:what emitInfraImports — EventBus/authz 인프라 import 출력

package ssac

import "strings"

// emitInfraImports writes EventBus and authz imports when needed.
func emitInfraImports(b *strings.Builder, d importData) {
	if d.HasPublish {
		b.WriteString("from app.dependencies.event_bus import EventBus\n")
	}
	if d.HasAuth {
		b.WriteString("from app.dependencies.authz import authz_check\n")
	}
}
