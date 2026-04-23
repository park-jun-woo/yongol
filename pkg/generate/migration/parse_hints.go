//ff:func feature=migration type=parser control=iteration dimension=1 topic=migration-hints
//ff:what ParseHints — DDL HintComment 리스트로부터 Hints 구조체 구성
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// ParseHints converts raw hint comments extracted from DDL files into
// the Hints structure consumed by Diff / check_safety. An empty list
// yields a non-nil but empty Hints.
func ParseHints(comments []ddl.HintComment) *Hints {
	h := newEmptyHints()
	for _, c := range comments {
		applyHintComment(h, c)
	}
	return h
}
