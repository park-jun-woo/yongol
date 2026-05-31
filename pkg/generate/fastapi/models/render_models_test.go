//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderModels — DDL Table → SQLAlchemy 모델 소스 생성(헤더·클래스·컬럼·FK·table_args) 검증
package models

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderModels(t *testing.T) {
	tables := []ddl.Table{
		{
			Name:        "users",
			ColumnOrder: []string{"id", "email", "bio", "missing"},
			Columns: map[string]ddl.Column{
				"id":    {Name: "id", RawType: "BIGINT", NotNull: true},
				"email": {Name: "email", RawType: "VARCHAR(255)", NotNull: true},
				"bio":   {Name: "bio", RawType: "TEXT", NotNull: false, HasDefault: true, DefaultLiteral: "n/a"},
				// "missing" intentionally absent from Columns -> skip branch
			},
			PrimaryKey: []string{"id"},
			Indexes: []ddl.Index{
				{Name: "uq_email", Columns: []string{"email"}, IsUnique: true},
			},
		},
		{
			Name:        "posts",
			ColumnOrder: []string{"id", "user_id"},
			Columns: map[string]ddl.Column{
				"id":      {Name: "id", RawType: "BIGINT", NotNull: true},
				"user_id": {Name: "user_id", RawType: "BIGINT", NotNull: true},
			},
			PrimaryKey:  []string{"id"},
			ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}},
		},
	}

	out, err := RenderModels(tables)
	if err != nil {
		t.Fatalf("RenderModels error: %v", err)
	}

	wantSubstrings := []string{
		"from __future__ import annotations",
		"from app.database import Base",
		"class User(Base):",
		`__tablename__ = "users"`,
		"id: Mapped[int] = mapped_column(Integer, primary_key=True)",
		"email: Mapped[str] = mapped_column(String)",
		`bio: Mapped[str | None] = mapped_column(Text, nullable=True, default="n/a")`,
		`UniqueConstraint("email", name="uq_email")`,
		"class Post(Base):",
		`user_id: Mapped[int] = mapped_column(Integer, ForeignKey("users.id"))`,
	}
	for _, w := range wantSubstrings {
		if !strings.Contains(out, w) {
			t.Errorf("RenderModels output missing %q\n---\n%s", w, out)
		}
	}

	// "missing" column was skipped, so it must not appear.
	if strings.Contains(out, "missing:") {
		t.Errorf("skipped column should not be rendered:\n%s", out)
	}
}
