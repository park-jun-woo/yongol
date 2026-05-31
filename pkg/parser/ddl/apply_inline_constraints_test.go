//ff:func feature=manifest type=test control=sequence
//ff:what applyInlineConstraints — 인라인 PK / UNIQUE / FK 반영
package ddl

import (
	"testing"
)

func TestApplyInlineConstraints(t *testing.T) {
	t.Run("primary key", func(t *testing.T) {
		tb := &Table{}
		applyInlineConstraints(tb, "ID BIGINT PRIMARY KEY", "id", []string{"id", "BIGINT", "PRIMARY", "KEY"})
		if len(tb.PrimaryKey) != 1 || tb.PrimaryKey[0] != "id" {
			t.Errorf("PrimaryKey = %v", tb.PrimaryKey)
		}
		if len(tb.Indexes) != 0 {
			t.Errorf("PRIMARY should not add unique index")
		}
	})
	t.Run("unique non-primary", func(t *testing.T) {
		tb := &Table{}
		applyInlineConstraints(tb, "EMAIL VARCHAR UNIQUE", "email", []string{"email", "VARCHAR", "UNIQUE"})
		if len(tb.Indexes) != 1 || tb.Indexes[0].Name != "email_unique" || !tb.Indexes[0].IsUnique {
			t.Errorf("Indexes = %+v", tb.Indexes)
		}
	})
	t.Run("inline references", func(t *testing.T) {
		tb := &Table{}
		applyInlineConstraints(tb, "ORG_ID BIGINT REFERENCES ORGS(ID)", "org_id",
			[]string{"org_id", "BIGINT", "REFERENCES", "orgs(id)"})
		if len(tb.ForeignKeys) != 1 {
			t.Fatalf("ForeignKeys = %+v", tb.ForeignKeys)
		}
		fk := tb.ForeignKeys[0]
		if fk.Column != "org_id" || fk.RefTable != "orgs" {
			t.Errorf("fk = %+v", fk)
		}
	})
}
