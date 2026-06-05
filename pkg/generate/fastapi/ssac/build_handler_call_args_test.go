//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestBuildHandlerCallArgs — buildHandlerCallArgs 호출 인자(session/path/body/query/user/event_bus) 구성 검증
package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestBuildHandlerCallArgs(t *testing.T) {
	plan := &ir.ServicePlan{
		PathParams:  []string{"id"},
		QueryParams: []ir.QueryParamMeta{{Name: "limit"}, {Name: "cursor"}},
	}
	cases := []struct {
		name                         string
		hasBody, isPreAuth, eventBus bool
		want                         []string
	}{
		{
			name:    "full args",
			want:    []string{"session", "id", "body", "limit", "cursor", "current_user", "event_bus"},
			hasBody: true, isPreAuth: false, eventBus: true,
		},
		{
			name:    "no body, pre-auth, no bus",
			want:    []string{"session", "id", "limit", "cursor"},
			hasBody: false, isPreAuth: true, eventBus: false,
		},
		{
			name:    "body only, no user/bus when preauth",
			want:    []string{"session", "id", "body", "limit", "cursor"},
			hasBody: true, isPreAuth: true, eventBus: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := buildHandlerCallArgs(plan, c.hasBody, c.isPreAuth, c.eventBus)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("buildHandlerCallArgs() = %v, want %v", got, c.want)
			}
		})
	}
}
