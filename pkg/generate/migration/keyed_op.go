//ff:type feature=migration type=model
//ff:what keyedOp — sortByDependency 의 phase+order 키드 래퍼
package migration

// keyedOp carries the phase / original order of an Operation for the
// dependency sort.
type keyedOp struct {
	phase int
	order int
	op    Operation
}
