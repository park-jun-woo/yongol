//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallErrorLines — @call 이후 err!=nil 분기 에러 응답 라인

package ssac

import "fmt"

func (g *methodGen) buildCallErrorLines(pkgName, callFunc, msg string, status int) []string {
	lines := []string{"if err != nil {"}
	if g.IsSubscribe {
		lines = append(lines, fmt.Sprintf("\treturn fmt.Errorf(\"%s.%s: %%w\", err)", pkgName, callFunc))
	} else {
		lines = append(lines,
			fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
			fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		)
	}
	lines = append(lines, "}")
	return lines
}
