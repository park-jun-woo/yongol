//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what formatCandidateNames — 후보 detectedAuthOp[] → OpID 문자열 slice

package hurl

// formatCandidateNames returns the operationId list for a set of auth
// candidates, preserving the caller's ordering. Kept separate so
// pickCandidate stays at control=sequence (no iteration at depth 1).
func formatCandidateNames(cands []detectedAuthOp) []string {
	names := make([]string, len(cands))
	for i, c := range cands {
		names[i] = c.OpID
	}
	return names
}
