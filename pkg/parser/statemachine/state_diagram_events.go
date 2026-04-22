//ff:func feature=statemachine type=util control=iteration dimension=1 topic=states
//ff:what 다이어그램의 모든 고유 이벤트 이름을 반환한다
package statemachine

// Events returns all unique event names in this diagram.
func (d *StateDiagram) Events() []string {
	seen := make(map[string]bool)
	var result []string
	for _, t := range d.Transitions {
		if !seen[t.Event] {
			seen[t.Event] = true
			result = append(result, t.Event)
		}
	}
	return result
}
