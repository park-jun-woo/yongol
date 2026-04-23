//ff:func feature=migration type=util control=sequence
//ff:what newSplitState — splitState 초기값 (빈 out / 0 depth)
package migration

// newSplitState returns an initialised splitState.
func newSplitState() *splitState {
	return &splitState{out: []string{}}
}
