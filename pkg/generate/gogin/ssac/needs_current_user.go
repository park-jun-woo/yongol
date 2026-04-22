//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what needsCurrentUser — SSaC 시퀀스 중 currentUser 참조가 존재하는지 검사

package ssac

import (
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// needsCurrentUser checks if any sequence references currentUser.
func needsCurrentUser(sf ssacparser.ServiceFunc) bool {
	for _, seq := range sf.Sequences {
		if sequenceUsesCurrentUser(seq) {
			return true
		}
	}
	return false
}
