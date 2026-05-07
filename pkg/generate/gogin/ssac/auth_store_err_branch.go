//ff:func feature=gen-gogin type=util control=selection
//ff:what authStoreErrBranch — error-handling branch for auth RefreshRotate/Logout call blocks

package ssac

import "fmt"

// authStoreErrBranch emits the `if err != nil { ... }` tail common to both
// auth RefreshRotate and Logout call blocks. Subscribe handlers bubble the
// wrapped error; HTTP handlers log + return the neutral JSON envelope.
func (g *methodGen) authStoreErrBranch(pkgName, callFunc string, status int, msg string) []string {
	lines := []string{"if err != nil {"}
	if g.IsSubscribe {
		lines = append(lines, fmt.Sprintf("\treturn fmt.Errorf(\"%s.%s: %%w\", err)", pkgName, callFunc))
	} else {
		lines = append(lines,
			fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
			fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: %q}, nil", g.FuncName, status, msg, neutralCode(status)),
		)
	}
	lines = append(lines, "}")
	return lines
}
