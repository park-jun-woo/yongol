//ff:func feature=migration type=test control=sequence
//ff:what TestEmitSQL_Determinism — 같은 입력·시각이면 결과가 동일
package migration

import (
	"testing"
	"time"
)

func TestEmitSQL_Determinism(t *testing.T) {
	ops := []Operation{
		AddColumn{Table: "t", Column: &Column{Name: "c", Type: CanonicalType{Base: "TEXT"}, Nullable: true}},
	}
	ts := time.Date(2026, 4, 22, 15, 30, 0, 0, time.UTC)
	a := EmitSQL(ops, EmitOptions{YongolVersion: "v1", GeneratedAt: ts})
	b := EmitSQL(ops, EmitOptions{YongolVersion: "v1", GeneratedAt: ts})
	if a != b {
		t.Errorf("non-deterministic output:\na=%s\nb=%s", a, b)
	}
}
