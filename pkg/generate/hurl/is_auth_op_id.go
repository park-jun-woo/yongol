//ff:func feature=gen-hurl type=util control=sequence
//ff:what isAuthOpID — operationID가 Register/Login인지 판정
package hurl

// isAuthOpID returns true if the operationID is a Register or Login operation.
func isAuthOpID(opID string) bool {
	return opID == "Register" || opID == "register" || opID == "Login" || opID == "login"
}
