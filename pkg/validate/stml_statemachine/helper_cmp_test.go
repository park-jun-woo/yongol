//ff:func feature=validate type=test-helper control=sequence dimension=1 topic=stml-statemachine
//ff:what cmp — model.status=value 비교 leaf GuardExpr를 생성하는 테스트 픽스처

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

func cmp(model, value string) *stml.GuardExpr {
	return &stml.GuardExpr{Kind: stml.GuardCompare, Ref: stml.GuardRef{Model: model, Field: "status"}, Value: value}
}
