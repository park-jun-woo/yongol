//ff:func feature=stml-gen type=util control=sequence
//ff:what operationID가 Login인지 판별한다
package stml

// isLoginAction returns true if the operationID is "Login".
func isLoginAction(operationID string) bool {
	return operationID == "Login"
}
