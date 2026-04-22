//ff:type feature=orchestrator type=test
//ff:what Level 상수 값 lock-in 테스트
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestLevelConstants 는 Level 상수 문자열 값이 의도된 값임을 고정한다.
// 값 변경 시 의도적으로 실패하므로, 변경이 필요하면 이 테스트도 함께 수정.
func TestLevelConstants(t *testing.T) {
	cases := []struct {
		name string
		got  diagnostic.Level
		want string
	}{
		{"LevelError", diagnostic.LevelError, "ERROR"},
		{"LevelWarning", diagnostic.LevelWarning, "WARNING"},
	}

	for _, c := range cases {
		if string(c.got) != c.want {
			t.Errorf("%s value changed: want %q, got %q", c.name, c.want, string(c.got))
		}
	}
}

// TestLevel_DistinctValues 는 Error / Warning 값이 서로 달라야 함을 확인한다.
func TestLevel_DistinctValues(t *testing.T) {
	if diagnostic.LevelError == diagnostic.LevelWarning {
		t.Errorf("LevelError and LevelWarning must differ, both=%q", diagnostic.LevelError)
	}
}

// TestLevel_TypeIsStringAlias 는 Level 이 string 기반 타입으로 변환 가능한지 확인한다.
func TestLevel_TypeIsStringAlias(t *testing.T) {
	l := diagnostic.LevelError
	if string(l) != "ERROR" {
		t.Errorf("string(LevelError): want %q, got %q", "ERROR", string(l))
	}

	custom := diagnostic.Level("INFO")
	if string(custom) != "INFO" {
		t.Errorf("Level(\"INFO\"): round-trip failed, got %q", string(custom))
	}
}
