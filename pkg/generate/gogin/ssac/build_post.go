//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildPost — @post 시퀀스 빌더 (sqlc INSERT :one)

package ssac

import ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"

func (g *methodGen) buildPost(seq ssacparser.Sequence) ([]string, []string) {
	// POST INSERT :one — not-found is never a normal flow; propagate as 500.
	return g.buildGet(seq, nil)
}
