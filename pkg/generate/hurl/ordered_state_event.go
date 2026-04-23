//ff:type feature=gen-hurl type=model
//ff:what orderedStateEvent — buildStateTransitions 가 소비하는 정렬된 이벤트 엔트리

package hurl

// orderedStateEvent is a single entry in the deduplicated, ordered list
// that drives smoke state-transition emission. ord mirrors the index
// assigned by buildTransitionOrder; event is the operationId that
// triggers the transition.
type orderedStateEvent struct {
	event string
	ord   int
}
