//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what 단일 Hurl 파일 라인을 파싱하여 요청/응답 상태 갱신
package hurl

// processHurlLine processes a single line from a .hurl file, updating the current entry state.
func processHurlLine(line string, lineNum int, path string, current *HurlEntry, entries []HurlEntry) (*HurlEntry, []HurlEntry) {
	if m := reHurlRequest.FindStringSubmatch(line); m != nil {
		if current != nil {
			entries = append(entries, *current)
		}
		current = &HurlEntry{
			Method: m[1],
			Path:   m[2],
			File:   path,
			Line:   lineNum,
		}
		return current, entries
	}

	if m := reHurlResponse.FindStringSubmatch(line); m != nil && current != nil {
		current.StatusCode = m[1]
	}

	return current, entries
}
