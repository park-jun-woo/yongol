//ff:func feature=gen-gogin type=util control=sequence
//ff:what returnErr — error 전파 return 문 (HTTP: 2-value, subscribe: single error)

package ssac

// returnErr returns the error-propagation return statement.
// HTTP handlers: "return nil, err" (2-value).
// Subscribe handlers: "return err" (single error).
func (g *methodGen) returnErr() string {
	if g.IsSubscribe {
		return "return err"
	}
	return "return nil, err"
}
