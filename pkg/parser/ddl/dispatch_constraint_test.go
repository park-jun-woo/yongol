//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what dispatchConstraint — 라인 종류별 분기 (FK / PRIMARY / UNIQUE / CHECK / column)

package ddl

import "testing"

func newTable() *Table {
	return &Table{Columns: map[string]Column{}}
}

func TestDispatchConstraint(t *testing.T) {
	t.Run("constraint foreign key", func(t *testing.T) {
		tb := newTable()
		line := "CONSTRAINT fk FOREIGN KEY (org_id) REFERENCES orgs(id)"
		dispatchConstraint(line, "CONSTRAINT FK FOREIGN KEY (ORG_ID) REFERENCES ORGS(ID)", tb, nil, false)
		if len(tb.ForeignKeys) != 1 || tb.ForeignKeys[0].RefTable != "orgs" {
			t.Errorf("ForeignKeys = %+v", tb.ForeignKeys)
		}
	})
	t.Run("primary key", func(t *testing.T) {
		tb := newTable()
		dispatchConstraint("PRIMARY KEY (id, org_id)", "PRIMARY KEY (ID, ORG_ID)", tb, nil, false)
		if len(tb.PrimaryKey) != 2 {
			t.Errorf("PrimaryKey = %v", tb.PrimaryKey)
		}
	})
	t.Run("unique", func(t *testing.T) {
		tb := newTable()
		dispatchConstraint("UNIQUE (email)", "UNIQUE (EMAIL)", tb, nil, false)
		if len(tb.Indexes) != 1 || !tb.Indexes[0].IsUnique {
			t.Errorf("Indexes = %+v", tb.Indexes)
		}
	})
	t.Run("table check enum", func(t *testing.T) {
		tb := newTable()
		tb.Columns["status"] = Column{}
		dispatchConstraint("CHECK (status IN ('a','b'))", "CHECK (STATUS IN ('A','B'))", tb, nil, false)
		if len(tb.Columns["status"].CheckEnum) != 2 {
			t.Errorf("CheckEnum = %v", tb.Columns["status"].CheckEnum)
		}
	})
	t.Run("constraint fk not parseable", func(t *testing.T) {
		tb := newTable()
		dispatchConstraint("CONSTRAINT name", "CONSTRAINT NAME", tb, nil, false)
		if len(tb.ForeignKeys) != 0 {
			t.Errorf("expected no FK, got %+v", tb.ForeignKeys)
		}
	})
	t.Run("column definition fallthrough", func(t *testing.T) {
		tb := newTable()
		dispatchConstraint("id BIGINT NOT NULL", "ID BIGINT NOT NULL", tb, nil, false)
		if _, ok := tb.Columns["id"]; !ok {
			t.Errorf("expected column id parsed, got %v", tb.Columns)
		}
	})
}
