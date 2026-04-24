//ff:func feature=gen-hurl type=util control=sequence
//ff:what pickCandidate — 다수 auth 후보 중 알파벳 우선 선택 + WARNING

package hurl

import (
	"fmt"
	"sort"
)

// pickCandidate returns the deterministic pick from a candidate list:
// alphabetical-first operationId. Emits a WARNING when more than one
// candidate exists (user should resolve ambiguity by renaming).
func pickCandidate(cands []detectedAuthOp, role string, warnings *[]string) *detectedAuthOp {
	if len(cands) == 0 {
		return nil
	}
	sort.SliceStable(cands, func(i, j int) bool {
		return cands[i].OpID < cands[j].OpID
	})
	if len(cands) > 1 {
		names := formatCandidateNames(cands)
		*warnings = append(*warnings, fmt.Sprintf(
			"detect_auth_ops: multiple %s candidates %v — using %q",
			role, names, cands[0].OpID))
	}
	c := cands[0]
	return &c
}
