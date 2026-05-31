//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestColumnAdapter_ZeroCov(t *testing.T) {
	a := columnAdapter{ddl.Column{RawType: "VARCHAR(50)", CheckEnum: []string{"a", "b"}}}
	if a.RawType() != "VARCHAR(50)" {
		t.Errorf("RawType() = %q", a.RawType())
	}
	if got := a.CheckEnum(); len(got) != 2 || got[0] != "a" {
		t.Errorf("CheckEnum() = %v", got)
	}
}
