//ff:type feature=gen-gogin type=model
//ff:what MainBlock — main.go 생성의 단위 블록 (조건 + imports + lines)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// MainBlock is the unit of main.go generation. Each block represents one
// initialization section (db, jwt, authz, queue, ...). The orchestrator
// collects active blocks, deduplicates imports, and assembles the body.
//
// Funcs holds top-level function declarations appended AFTER main()'s closing
// brace. Used by env helper blocks (envInt, envDuration, envStringList,
// envBool) that must live at package scope. Unlike Lines, Funcs are raw
// declarations (each entry is a full `func name(...) T { ... }` block) and
// are written verbatim with blank-line separation between blocks.
type MainBlock struct {
	Name    string
	Active  func(fs *yongol.Fullstack) bool
	Imports []string
	Lines   []string
	Funcs   []string
}
