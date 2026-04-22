//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildDelete — @delete 시퀀스 빌더 (sqlc DELETE :exec)

package ssac

import ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"

func (g *methodGen) buildDelete(seq ssacparser.Sequence) ([]string, []string) {
	return g.buildPut(seq)
}
