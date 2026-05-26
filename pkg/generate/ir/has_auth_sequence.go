//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what hasAuthSequence -- SSaC ServiceFuncs 에 @auth 시퀀스 존재 여부

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasAuthSequence returns true when any SSaC function uses @auth.
func hasAuthSequence(fs *yongol.Fullstack) bool {
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "auth" {
				return true
			}
		}
	}
	return false
}
