//ff:func feature=crosscheck type=parser control=sequence topic=scenario-check
//ff:what parseState.flushRequestBody — 축적된 body buffer 를 parse 해 BodyFields 갱신

package hurl

import "strings"

// flushRequestBody parses the accumulated request-body buffer into the
// current entry's BodyFields if the buffer looks like a JSON object.
// Called when the HTTP status line terminates the request or when the
// entry is flushed at EOF.
func (s *parseState) flushRequestBody() {
	if s.current == nil {
		return
	}
	body := strings.TrimSpace(s.bodyBuf.String())
	s.bodyBuf.Reset()
	if body == "" {
		return
	}
	fields := extractJSONFieldNames(body)
	if len(fields) > 0 {
		s.current.BodyFields = append(s.current.BodyFields, fields...)
	}
}
