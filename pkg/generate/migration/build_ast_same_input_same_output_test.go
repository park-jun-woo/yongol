//ff:func feature=migration type=test control=sequence
//ff:what TestBuildAST_SameInputSameOutput — 같은 DDL 두 번 파싱해 동일 AST 산출 (결정성)
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
