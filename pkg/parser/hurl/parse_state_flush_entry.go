//ff:func feature=crosscheck type=parser control=sequence topic=scenario-check
//ff:what parseState.flushEntry — 현재 entry 를 entries 에 append 하고 상태 초기화

package hurl

// flushEntry appends the current entry to entries and resets state for
// the next request. Called when a new request line appears or at EOF.
func (s *parseState) flushEntry() {
	if s.current == nil {
		return
	}
	s.flushRequestBody()
	s.entries = append(s.entries, *s.current)
	s.current = nil
	s.section = ""
	s.bodyBuf.Reset()
}
