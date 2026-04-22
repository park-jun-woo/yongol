//ff:type feature=orchestrator type=test
//ff:what Phase 상수 값 lock-in 테스트
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestPhaseConstants 는 Phase 상수 문자열 값이 의도된 값임을 고정한다.
// 값 변경 시 의도적으로 실패하므로, 변경이 필요하면 이 테스트도 함께 수정.
func TestPhaseConstants(t *testing.T) {
	cases := []struct {
		name string
		got  diagnostic.Phase
		want string
	}{
		{"PhaseParse", diagnostic.PhaseParse, "parse"},
		{"PhaseValidate", diagnostic.PhaseValidate, "validate"},
	}

	for _, c := range cases {
		if string(c.got) != c.want {
			t.Errorf("%s value changed: want %q, got %q", c.name, c.want, string(c.got))
		}
	}
}

// TestPhase_TypeIsStringAlias 는 Phase 가 string 기반 타입으로 변환 가능한지 확인한다.
func TestPhase_TypeIsStringAlias(t *testing.T) {
	p := diagnostic.PhaseValidate
	s := string(p)
	if s != "validate" {
		t.Errorf("string(PhaseValidate): want %q, got %q", "validate", s)
	}

	// 반대 방향: string → Phase 로 캐스팅 가능해야 함.
	custom := diagnostic.Phase("custom")
	if string(custom) != "custom" {
		t.Errorf("Phase(\"custom\"): round-trip failed, got %q", string(custom))
	}
}
