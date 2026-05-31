//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockAuthInit — auth.Configure + RefreshStore 주입 (라우트 마운트 없음)
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockAuthInit_Absent(t *testing.T) {
	block := blockAuthInit(prepared.Auth{Present: false}, "example.com/zenflow")
	if block.Active == nil || block.Active(nil) {
		t.Errorf("absent auth must carry authAlwaysInactive predicate")
	}
	if len(block.Lines) != 0 {
		t.Errorf("absent auth must emit no lines, got %v", block.Lines)
	}
}
