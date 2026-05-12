//ff:func feature=stml-gen type=util control=sequence
//ff:what operationId가 DELETE 동작인지 판별한다
package stml

import "strings"

// isDeleteOperation returns true if the operationId represents a DELETE action.
func isDeleteOperation(operationID string) bool {
	return strings.HasPrefix(operationID, "Delete")
}
