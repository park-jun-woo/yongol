//ff:func feature=gen-hurl type=test-helper control=iteration dimension=1
//ff:what stepOpIDs — []step 에서 OperationID 필드를 순서대로 []string 으로 추출

package hurl

// stepOpIDs extracts OperationID from a slice of step in order. Shared
// across Phase003/004 smoke ordering tests.
func stepOpIDs(steps []step) []string {
	out := make([]string, 0, len(steps))
	for _, s := range steps {
		out = append(out, s.OperationID)
	}
	return out
}
