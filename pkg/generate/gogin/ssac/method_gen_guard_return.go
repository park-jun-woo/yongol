//ff:func feature=gen-gogin type=util control=sequence
//ff:what guardReturn — guard early-return 응답 라인 (HTTP: api.Response, subscribe: fmt.Errorf)

package ssac

import "fmt"

// guardReturn returns the guard (early-return) response line.
// HTTP handlers: "return api.<Op><Status>JSONResponse{Error: ..., Code: ...}, nil"
// Subscribe handlers: "return fmt.Errorf(<msg>)"
func (g *methodGen) guardReturn(msg string, status int) string {
	if g.IsSubscribe {
		return fmt.Sprintf("return fmt.Errorf(%q)", msg)
	}
	return fmt.Sprintf("return %s, nil", g.errorResponseLiteral(status, msg, neutralCode(status)))
}
