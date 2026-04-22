//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what assembleBody — 활성 블록의 Lines 를 순서대로 연결

package boot

// assembleBody concatenates Lines from all blocks with blank-line separation.
func assembleBody(blocks []MainBlock) []string {
	var body []string
	for i, b := range blocks {
		if i > 0 && len(b.Lines) > 0 {
			body = append(body, "")
		}
		body = append(body, b.Lines...)
	}
	return body
}
