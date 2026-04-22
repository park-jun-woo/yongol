//ff:func feature=gen-filefunc type=util control=selection
//ff:what resolveFeatureDescription — SSOT 설명 → 인프라 baseline → fallback 순으로 feature 설명 선택
package filefunc

// infraDescriptions holds fixed one-liners for infrastructure packages that
// every Go+Gin backend has, regardless of SSOT content.
var infraDescriptions = map[string]string{
	"api":          "OpenAPI 기반 Gin 라우터",
	"middleware":   "인증·CORS·request 검증",
	"model":        "DTO 구조체",
	"auth":         "JWT 발급·검증",
	"db":           "sqlc 생성 쿼리",
	"boot":         "main.go 및 초기화 블록",
	"service":      "비즈니스 로직 서비스 (SSaC @func 구현)",
	"statemachine": "상태 전이 테이블",
	"queue":        "큐 발행·구독 어댑터",
	"session":      "세션 백엔드 어댑터",
	"cache":        "캐시 백엔드 어댑터",
	"file":         "파일 스토리지 어댑터",
	"authz":        "OPA 기반 인가 검사",
	"config":       "환경 변수·런타임 설정",
	"dashboard":    "집계·리포트",
	"report":       "실행 리포트",
	"schedule":     "cron 파서",
	"resolver":     "계층 구조 해석",
	"billing":      "크레딧 차감·조회",
	"webhook":      "웹훅 구독·이벤트 발행",
	"audit":        "감사 로그",
	"org":          "조직 관리",
	"user":         "사용자 관리",
	"workflow":     "워크플로우 실행·복제·실행 시나리오",
	"template":     "템플릿 CRUD 및 마켓플레이스",
	"execution":    "실행 기록 조회",
}

// fallbackDescription is used when neither SSOT metadata nor infra baseline
// provides a description for a package.
const fallbackDescription = "백엔드 패키지"

// resolveFeatureDescription picks the best one-line description for a
// feature using SSOT → infra → fallback priority.
func resolveFeatureDescription(name, ssotDesc string) string {
	switch {
	case ssotDesc != "":
		return ssotDesc
	case hasInfraDescription(name):
		return infraDescriptions[name]
	default:
		return fallbackDescription
	}
}
