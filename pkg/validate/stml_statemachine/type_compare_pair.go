//ff:type feature=validate type=model topic=stml-statemachine
//ff:what comparePair — 가드 비교식 <model>.<field> = <value>에서 추출한 (model, value) 쌍

package stml_statemachine

// comparePair is the (model, value) pair extracted from a guard comparison such
// as "workflow.status = active" (Model="workflow", Value="active").
type comparePair struct {
	Model string
	Value string
}
