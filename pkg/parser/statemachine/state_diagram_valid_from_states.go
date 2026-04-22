//ff:func feature=statemachine type=util control=iteration dimension=1 topic=states
//ff:what 특정 이벤트가 유효한 출발 상태 목록을 반환한다
package statemachine

// ValidFromStates returns all states from which the given event is valid.
func (d *StateDiagram) ValidFromStates(event string) []string {
	var result []string
	for _, t := range d.Transitions {
		if t.Event == event {
			result = append(result, t.From)
		}
	}
	return result
}
