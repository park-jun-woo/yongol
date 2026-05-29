//ff:func feature=crosscheck type=parser control=sequence topic=scenario-check
//ff:what newRequestEntry — request 라인 매치 결과로부터 HurlEntry 생성

package hurl

// newRequestEntry builds a HurlEntry from a matched request-line regexp
// submatch (m[1]=method, m[2]=URL {{var}} or "", m[3]=path) at the given
// file/line position.
func newRequestEntry(m []string, file string, line int) *HurlEntry {
	return &HurlEntry{
		Method: m[1],
		Path:   trimQuery(m[3]),
		URLVar: m[2],
		File:   file,
		Line:   line,
	}
}
