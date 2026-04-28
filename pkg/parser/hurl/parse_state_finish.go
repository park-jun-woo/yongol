//ff:func feature=crosscheck type=parser control=sequence topic=scenario-check
//ff:what parseState.finish — EOF 시점에 마지막 entry 를 flush

package hurl

func (s *parseState) finish() {
	s.flushEntry()
}
