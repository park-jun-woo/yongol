//ff:func feature=stml-gen type=util control=sequence
//ff:what ActionBlock이 선언 기반 onSuccess(캡처 커밋·리다이렉트) 경로를 타는지 판별한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// actionHasFlowSuccess reports whether the action's onSuccess is driven by
// STML flow declarations (data-capture and/or data-redirect). When false the
// mutation keeps the default invalidateQueries() path.
func actionHasFlowSuccess(a stmlparser.ActionBlock, bearerAuth bool) bool {
	return len(actionFlowCaptures(a, bearerAuth)) > 0 || a.Redirect != ""
}
