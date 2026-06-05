//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-statemachine
//ff:what stateInSlice — 상태 이름 슬라이스에 특정 상태가 포함되는지 검사

package stml_statemachine

// stateInSlice reports whether states contains target.
func stateInSlice(states []string, target string) bool {
	for _, s := range states {
		if s == target {
			return true
		}
	}
	return false
}
