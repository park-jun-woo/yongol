//ff:func feature=gen-filefunc type=util control=sequence
//ff:what buildTopicEntries — 고정 topic 맵 반환 (Phase001 단계에서는 동적 확장 없음)
package filefunc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildTopicEntries returns the topic catalogue. For Phase001 we emit a fixed
// baseline set that covers the semantic units produced by SSaC @call /
// @publish / @subscribe / @state. fs is accepted for future dynamic expansion
// but currently unused.
func buildTopicEntries(fs *yongol.Fullstack) map[string]string {
	_ = fs
	return map[string]string{
		"request-binding":      "요청 파라미터 수집, body 바인딩",
		"response-serialize":   "응답 필드 구성",
		"transaction-boundary": "BeginTx / Commit / Rollback",
		"state-transition":     "@state 전이 실행",
		"auth-check":           "@auth 게이트",
		"pagination":           "페이지네이션 적용",
		"error-mapping":        "validation / domain / infra error 분류",
		"observability":        "slog / metric / trace",
		"publish":              "큐 발행",
		"subscribe":            "큐 구독",
		"pointer-helper":       "ptr/deref 제네릭 헬퍼",
	}
}
