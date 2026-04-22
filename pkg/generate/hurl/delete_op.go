//ff:type feature=gen-hurl type=model
//ff:what deleteOp — DELETE endpoint의 path/operationID/resource 정보
package hurl

// deleteOp holds info about a DELETE operation for sorting.
type deleteOp struct {
	path     string
	opID     string
	resource string
}
