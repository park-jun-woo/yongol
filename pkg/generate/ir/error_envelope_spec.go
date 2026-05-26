//ff:type feature=gen-ir type=model
//ff:what ErrorEnvelopeSpec -- 백엔드 간 동일한 에러 응답 JSON 구조 명세

package ir

// ErrorEnvelopeSpec defines the standard error response JSON structure
// that all backend renderers must produce. This ensures NestJS, FastAPI,
// and Go+Gin backends return identical error shapes to clients.
//
// Example: Fields = ["error", "message", "request_id"] produces:
//
//	{"error": "NOT_FOUND", "message": "...", "request_id": "01HZ..."}
type ErrorEnvelopeSpec struct {
	// Fields lists the JSON keys present in every error response body.
	Fields []string
}
