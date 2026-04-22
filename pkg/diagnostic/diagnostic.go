//ff:type feature=orchestrator type=model
//ff:what 파싱/검증/교차검증 단계의 진단 메시지
package diagnostic

// Diagnostic represents a single diagnostic message from any phase.
type Diagnostic struct {
	File    string // source file path
	Line    int    // line number (0 if unknown)
	Phase   Phase  // parse, validate
	Level   Level  // error, warning
	Message string // 본문 메시지 (Rule-ID + 무엇이 잘못)
	Advice  string // → 권고: 본문 (어떻게 고치는지). 비어있으면 미표시.
}
