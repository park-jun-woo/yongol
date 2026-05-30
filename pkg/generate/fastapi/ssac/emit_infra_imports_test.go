//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestEmitInfraImports — EventBus/authz import 조건부 출력

package ssac

import (
	"strings"
	"testing"
)

func TestEmitInfraImports(t *testing.T) {
	cases := []struct {
		name       string
		hasPublish bool
		hasAuth    bool
		want       string
	}{
		{"None", false, false, ""},
		{"PublishOnly", true, false, "from app.dependencies.event_bus import EventBus\n"},
		{"AuthOnly", false, true, "from app.dependencies.authz import authz_check\n"},
		{"Both", true, true, "from app.dependencies.event_bus import EventBus\nfrom app.dependencies.authz import authz_check\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var b strings.Builder
			emitInfraImports(&b, importData{HasPublish: c.hasPublish, HasAuth: c.hasAuth})
			if b.String() != c.want {
				t.Errorf("got %q, want %q", b.String(), c.want)
			}
		})
	}
}
