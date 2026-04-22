//ff:func feature=gen-filefunc type=util control=sequence
//ff:what buildTypeEntries — Go+Gin 백엔드에 등장하는 고정 type 카테고리 맵 반환
package filefunc

// buildTypeEntries returns the fixed type catalogue for generated Go+Gin
// backend code. Keys must stay in sync with //ff:type values emitted by the
// code generator (Phase002+).
func buildTypeEntries() map[string]string {
	return map[string]string{
		"handler":     "HTTP 요청 핸들러 (Gin handler)",
		"service":     "비즈니스 로직 함수 (SSaC @func)",
		"model":       "데이터 전송 객체 (DTO)",
		"query":       "sqlc 쿼리 래퍼",
		"middleware":  "Gin 미들웨어",
		"config":      "환경 변수·런타임 설정",
		"accessor":    "getter/setter/reexport",
		"util":        "유틸리티 함수",
		"generator":   "코드 생성 보조",
		"loader":      "초기화/로더",
		"command":     "엔트리포인트 함수",
		"test":        "테스트 함수",
		"test-helper": "테스트 헬퍼",
	}
}
