//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-openapi
//ff:what 액션 목록에서 auth.refresh 캡처가 존재하는지 판정한다

package stml_openapi

import (
	"slices"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func actionsHaveRefreshCapture(actions []stml.ActionBlock) bool {
	for _, a := range actions {
		if slices.ContainsFunc(a.Captures, func(c stml.CaptureBind) bool { return c.Sink == "auth.refresh" }) {
			return true
		}
	}
	return false
}
