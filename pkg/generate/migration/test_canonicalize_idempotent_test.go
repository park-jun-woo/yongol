//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Canonicalize idempotence — 같은 DDL 두 번 파싱하면 동일 AST
package migration

import (
	"reflect"
	"testing"
)

func TestBuildAST_SameInputSameOutput(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(32) NOT NULL DEFAULT 'member',
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
CREATE INDEX idx_users_role ON users (role);
`
	a := NewSchema()
	b := NewSchema()
	if err := BuildASTFromSQL(a, sql); err != nil {
		t.Fatal(err)
	}
	if err := BuildASTFromSQL(b, sql); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("determinism broken: %+v vs %+v", a.Tables["users"], b.Tables["users"])
	}
}

func TestNormalizeType_SameInputSameOutput(t *testing.T) {
	// int / int4 / integer / INTEGER must collapse to the same canonical.
	ct1, _ := NormalizeType("int")
	ct2, _ := NormalizeType("int4")
	ct3, _ := NormalizeType("INTEGER")
	if !ct1.Equal(ct2) || !ct2.Equal(ct3) {
		t.Errorf("INTEGER aliases not equal: %+v %+v %+v", ct1, ct2, ct3)
	}
	// timestamp variants.
	ts1, _ := NormalizeType("timestamptz")
	ts2, _ := NormalizeType("timestamp with time zone")
	if !ts1.Equal(ts2) {
		t.Errorf("TIMESTAMPTZ aliases not equal: %+v %+v", ts1, ts2)
	}
}
