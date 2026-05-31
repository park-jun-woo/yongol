//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"
)

func TestHeadPredicates_ZeroCov(t *testing.T) {
	if !isBooleanHead("BOOLEAN") || isBooleanHead("TEXT") {
		t.Errorf("isBooleanHead mismatch")
	}
	if !isFloatHead("REAL") || isFloatHead("TEXT") {
		t.Errorf("isFloatHead mismatch")
	}
	if !isIntegerHead("BIGINT") || isIntegerHead("TEXT") {
		t.Errorf("isIntegerHead mismatch")
	}
	if !isStringHead("TEXT") || isStringHead("BIGINT") {
		t.Errorf("isStringHead mismatch")
	}
}
