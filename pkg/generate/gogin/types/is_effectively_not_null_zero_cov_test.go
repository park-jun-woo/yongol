//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsEffectivelyNotNull_ZeroCov(t *testing.T) {
	if isEffectivelyNotNull(ddl.Column{NullableAnnot: true, NotNull: true}) {
		t.Errorf("@nullable annotation must override NOT NULL")
	}
	if !isEffectivelyNotNull(ddl.Column{NotNull: true}) {
		t.Errorf("NOT NULL without annotation must be true")
	}
	if isEffectivelyNotNull(ddl.Column{NotNull: false}) {
		t.Errorf("nullable column must be false")
	}
}
