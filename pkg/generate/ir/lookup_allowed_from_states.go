//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what lookupAllowedFromStates -- diagramID/transition 에 매칭되는 stateDiagram 의 허용 source 상태 조회

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// lookupAllowedFromStates returns the valid source states for the given
// transition by locating the matching state diagram (by ID or Symbol).
// Returns nil when no diagram matches.
func lookupAllowedFromStates(fs *yongol.Fullstack, diagramID, transition string) []string {
	for _, d := range fs.StateDiagrams {
		if d.ID == diagramID || d.Symbol == diagramID {
			return d.ValidFromStates(transition)
		}
	}
	return nil
}
