//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-statemachine
//ff:what pageWithAction — 단일 ActionBlock을 담은 테스트용 STML PageSpec 생성

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

func pageWithAction(a stml.ActionBlock) stml.PageSpec {
	return stml.PageSpec{FileName: "workflow-page.html", Actions: []stml.ActionBlock{a}}
}
