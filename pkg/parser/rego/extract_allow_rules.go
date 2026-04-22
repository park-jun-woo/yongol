//ff:func feature=policy type=parser control=iteration dimension=1
//ff:what extractAllowRules — Rego 소스에서 allow 블록들을 분리하고 AllowRule 추출
package rego

import "strings"

func extractAllowRules(content string, p *Policy) {
	normalized := strings.ReplaceAll(content, "\nallow {", "\nallow if {")
	// 원본 content 기준 라인 번호 계산을 위해 개행 누적 수를 추적.
	// "\nallow {" → "\nallow if {" 치환은 동일 줄 내 토큰만 늘리므로 줄 번호는 보존.
	const token = "allow if {"
	offset := 0
	for {
		idx := strings.Index(normalized[offset:], token)
		if idx < 0 {
			break
		}
		headerStart := offset + idx
		bodyStart := headerStart + len(token)
		rel := findClosingBrace(normalized[bodyStart:])
		if rel < 0 {
			offset = bodyStart
			continue
		}
		endIdx := bodyStart + rel
		block := normalized[bodyStart:endIdx]
		if rule, ok := processAllowBlock(block); ok {
			rule.SourceLine = lineOfOffset(normalized, headerStart)
			p.Rules = append(p.Rules, rule)
		}
		offset = endIdx + 1
	}
}

// lineOfOffset — 문자열 s 의 바이트 오프셋 off 가 놓인 1-based 라인 번호.
// off 가 범위 밖이면 0 을 반환한다.
func lineOfOffset(s string, off int) int {
	if off < 0 || off > len(s) {
		return 0
	}
	return 1 + strings.Count(s[:off], "\n")
}
