//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildPut — @put 시퀀스 빌더 (sqlc UPDATE :exec)

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildPut(seq ssacparser.Sequence) ([]string, []string) {
	method := resolveSQLCMethod(seq.Model)
	assign := g.assignOp(false)
	argStr := g.sqlcArgs(method, seq.Inputs)
	return []string{
		fmt.Sprintf("err %s %s.%s(%s)", assign, g.queryVar(), method, argStr),
		"if err != nil { return nil, err }",
	}, nil
}
