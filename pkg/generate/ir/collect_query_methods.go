//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what collectQueryMethods -- SSaC CRUD 시퀀스에서 sqlc QueryMethod 목록 추출

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectQueryMethods extracts the QueryMethod list from SSaC sequences that
// reference sqlc queries (@get/@post/@put/@delete).
func collectQueryMethods(seqs []ssac.Sequence) []QueryMethod {
	var methods []QueryMethod
	seen := make(map[string]bool)

	for _, seq := range seqs {
		if !isCRUDSeq(seq.Type) {
			continue
		}
		fullMethod := strings.ReplaceAll(seq.Model, ".", "")
		if seen[fullMethod] {
			continue
		}
		seen[fullMethod] = true
		methods = append(methods, QueryMethod{
			Name:    fullMethod,
			Package: seq.Package,
		})
	}
	return methods
}
