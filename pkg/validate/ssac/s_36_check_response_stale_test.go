//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what s36CheckResponseStale — stale 변수 참조 검출 검증 (stale 참조 → WARNING, 비 stale 통과, 비 response 스킵)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS36CheckResponseStale(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fn := parsessac.ServiceFunc{FileName: "order.ssac"}
		seq := parsessac.Sequence{Type: "response", Line: 10, Fields: map[string]string{"data": "order.Name"}}
		stale := map[string]bool{"order": true}
		diags := s36CheckResponseStale(fn, 0, seq, stale)
		if len(diags) != 1 {
			t.Fatalf("got %d diags, want 1", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[S-36]") {
			t.Errorf("Message = %q, want [S-36]", diags[0].Message)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %q, want WARNING", diags[0].Level)
		}
	})
	t.Run("NonResponseSkipped", func(t *testing.T) {
		fn := parsessac.ServiceFunc{FileName: "order.ssac"}
		seq := parsessac.Sequence{Type: "get", Line: 10}
		stale := map[string]bool{"order": true}
		diags := s36CheckResponseStale(fn, 0, seq, stale)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("NotStalePasses", func(t *testing.T) {
		fn := parsessac.ServiceFunc{FileName: "order.ssac"}
		seq := parsessac.Sequence{Type: "response", Line: 10, Fields: map[string]string{"data": "order.Name"}}
		stale := map[string]bool{"order": false}
		diags := s36CheckResponseStale(fn, 0, seq, stale)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
	t.Run("SuppressWarn", func(t *testing.T) {
		fn := parsessac.ServiceFunc{FileName: "order.ssac"}
		seq := parsessac.Sequence{Type: "response", Line: 10, Fields: map[string]string{"data": "order.Name"}, SuppressWarn: true}
		stale := map[string]bool{"order": true}
		diags := s36CheckResponseStale(fn, 0, seq, stale)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 (SuppressWarn should skip)", len(diags))
		}
	})
}
