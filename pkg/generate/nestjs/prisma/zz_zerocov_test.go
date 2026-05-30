package prisma

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증

func TestIsPrimaryKey_ZeroCov(t *testing.T) {
	if !isPrimaryKey("id", []string{"id", "tenant_id"}) {
		t.Error("expected id to be primary key")
	}
	if isPrimaryKey("name", []string{"id"}) {
		t.Error("expected name not to be primary key")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsColumnUnique_ZeroCov — 단일 컬럼 unique 인덱스 매칭/비매칭

func TestIsColumnUnique_ZeroCov(t *testing.T) {
	idxs := []ddl.Index{
		{Name: "u1", Columns: []string{"email"}, IsUnique: true},
		{Name: "u2", Columns: []string{"a", "b"}, IsUnique: true},
		{Name: "n1", Columns: []string{"name"}, IsUnique: false},
	}
	if !isColumnUnique("email", idxs) {
		t.Error("expected email unique")
	}
	if isColumnUnique("a", idxs) {
		t.Error("composite index should not count as single-column unique")
	}
	if isColumnUnique("name", idxs) {
		t.Error("non-unique index should not count")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestPascalCase_ZeroCov — snake_case → PascalCase (빈 파트 포함)

func TestPascalCase_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"user_account": "UserAccount",
		"post":         "Post",
		"a__b":         "AB",
		"":             "",
	}
	for in, want := range cases {
		if got := pascalCase(in); got != want {
			t.Errorf("pascalCase(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestSingularize_ZeroCov — 복수형 접미사 각 분기

func TestSingularize_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"categories": "category",
		"classes":    "class",
		"boxes":      "box",
		"users":      "user",
		"address":    "address", // ss → unchanged by last branch
		"item":       "item",
	}
	for in, want := range cases {
		if got := singularize(in); got != want {
			t.Errorf("singularize(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestMapPGFamily_ZeroCov — 모든 PG 패밀리 → Prisma 스칼라 매핑

func TestMapPGFamily_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"BIGINT": "Int", "INTEGER": "Int", "SMALLINT": "Int",
		"TEXT": "String", "CITEXT": "String",
		"BOOLEAN": "Boolean", "BOOL": "Boolean",
		"TIMESTAMP": "DateTime", "DATE": "DateTime",
		"UUID": "String",
		"JSONB": "Json", "JSON": "Json",
		"NUMERIC": "Decimal", "DECIMAL": "Decimal",
		"FLOAT": "Float", "REAL": "Float",
		"BYTEA": "Bytes",
		"INET": "String", "INTERVAL": "String",
		"UNKNOWNTYPE": "String",
	}
	for in, want := range cases {
		if got := mapPGFamily(in); got != want {
			t.Errorf("mapPGFamily(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestPgToPrismaType_ZeroCov — array/precision 정리 후 매핑

func TestPgToPrismaType_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"BIGINT":       "Int",
		"varchar(255)": "String",
		"text[]":       "String[]",
		"numeric(10,2)": "Decimal",
	}
	for in, want := range cases {
		if got := pgToPrismaType(in); got != want {
			t.Errorf("pgToPrismaType(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestPrismaDefault_ZeroCov — 기본값 리터럴 변환 분기

func TestPrismaDefault_ZeroCov(t *testing.T) {
	cases := []struct {
		lit  string
		want string
	}{
		{"", `""`},
		{"now()", "now()"},
		{"CURRENT_TIMESTAMP", "now()"},
		{"gen_random_uuid()", "uuid()"},
		{"uuid_generate_v4()", "uuid()"},
		{"TRUE", "true"},
		{"false", "false"},
		{"42", "42"},
		{"3.14", "3.14"},
		{"hello", `"hello"`},
	}
	for _, c := range cases {
		col := ddl.Column{DefaultLiteral: c.lit}
		if got := prismaDefault(col); got != c.want {
			t.Errorf("prismaDefault(%q)=%q want %q", c.lit, got, c.want)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestColumnAttributes_ZeroCov — PK/UUID/timestamp/serial/identity/default 분기

func TestColumnAttributes_ZeroCov(t *testing.T) {
	pk := []string{"id"}
	if got := columnAttributes(ddl.Column{RawType: "BIGINT"}, "id", pk); got != "@id" {
		t.Errorf("pk attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "UUID", HasDefault: true}, "uid", pk); got != "@default(uuid())" {
		t.Errorf("uuid attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TIMESTAMPTZ", HasDefault: true}, "ts", pk); got != "@default(now())" {
		t.Errorf("ts attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "BIGSERIAL"}, "n", pk); got != "@default(autoincrement())" {
		t.Errorf("serial attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "BIGINT", IsIdentity: true}, "n", pk); got != "@default(autoincrement())" {
		t.Errorf("identity attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TEXT", HasDefault: true, DefaultLiteral: "x"}, "n", pk); got != `@default("x")` {
		t.Errorf("default attr = %q", got)
	}
	if got := columnAttributes(ddl.Column{RawType: "TEXT"}, "n", pk); got != "" {
		t.Errorf("no attr expected, got %q", got)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderColumn_ZeroCov — optional/unique/attr 조합 렌더링

func TestRenderColumn_ZeroCov(t *testing.T) {
	var b strings.Builder
	// PK column (not optional, @id)
	renderColumn(&b, ddl.Column{RawType: "BIGINT", NotNull: true}, "id", []string{"id"})
	// nullable column, no index → optional "?"
	renderColumn(&b, ddl.Column{RawType: "TEXT"}, "bio", []string{"id"})
	// unique column with attrs
	idxs := []ddl.Index{{Columns: []string{"email"}, IsUnique: true}}
	renderColumn(&b, ddl.Column{RawType: "TEXT", NotNull: true}, "email", []string{"id"}, idxs)
	// unique column without other attrs
	idxs2 := []ddl.Index{{Columns: []string{"slug"}, IsUnique: true}}
	renderColumn(&b, ddl.Column{RawType: "TEXT", NotNull: true}, "slug", []string{"id"}, idxs2)
	out := b.String()
	if !strings.Contains(out, "@id") {
		t.Error("expected @id")
	}
	if !strings.Contains(out, "String?") {
		t.Error("expected optional String?")
	}
	if !strings.Contains(out, "@unique") {
		t.Error("expected @unique")
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestDedupReverseRelations_ZeroCov — 중복 제거

func TestDedupReverseRelations_ZeroCov(t *testing.T) {
	rels := []reverseRelation{
		{FieldName: "posts", ModelName: "Post"},
		{FieldName: "posts", ModelName: "Post"},
		{FieldName: "comments", ModelName: "Comment"},
	}
	got := dedupReverseRelations(rels)
	if len(got) != 2 {
		t.Errorf("expected 2 deduped, got %d", len(got))
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestUniqueReverseFieldName_ZeroCov — base 그대로 반환

func TestUniqueReverseFieldName_ZeroCov(t *testing.T) {
	if got := uniqueReverseFieldName("posts", 3); got != "posts" {
		t.Errorf("got %q", got)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildReverseRelations_ZeroCov — FK 스캔 → 역관계 맵

func TestBuildReverseRelations_ZeroCov(t *testing.T) {
	tables := []ddl.Table{
		{Name: "users"},
		{Name: "posts", ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}}},
	}
	rm := buildReverseRelations(tables)
	revs := rm["users"]
	if len(revs) != 1 || revs[0].ModelName != "Post" || revs[0].FieldName != "posts" {
		t.Errorf("unexpected reverse relations: %+v", revs)
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderModel_ZeroCov — 컬럼/FK/역관계/인덱스/@@map 렌더링

func TestRenderModel_ZeroCov(t *testing.T) {
	var b strings.Builder
	table := ddl.Table{
		Name:        "posts",
		ColumnOrder: []string{"id", "user_id", "title", "missing"},
		Columns: map[string]ddl.Column{
			"id":      {RawType: "BIGINT", NotNull: true},
			"user_id": {RawType: "BIGINT", NotNull: true},
			"title":   {RawType: "TEXT"},
		},
		PrimaryKey:  []string{"id"},
		ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}},
		Indexes: []ddl.Index{
			{Columns: []string{"title"}, IsUnique: false},
			{Columns: []string{"user_id", "title"}, IsUnique: true},
		},
	}
	revRels := []reverseRelation{{FieldName: "comments", ModelName: "Comment"}}
	if err := renderModel(&b, table, revRels); err != nil {
		t.Fatalf("renderModel error: %v", err)
	}
	out := b.String()
	for _, want := range []string{"model Post {", "user", "User @relation", "comments", "Comment[]", "@@index([title])", "@@unique([user_id, title])", `@@map("posts")`} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n%s", want, out)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderSchema_ZeroCov — 전체 스키마 생성 (datasource/generator/model)

func TestRenderSchema_ZeroCov(t *testing.T) {
	tables := []ddl.Table{
		{
			Name:        "users",
			ColumnOrder: []string{"id"},
			Columns:     map[string]ddl.Column{"id": {RawType: "BIGINT", NotNull: true}},
			PrimaryKey:  []string{"id"},
		},
		{
			Name:        "posts",
			ColumnOrder: []string{"id", "user_id"},
			Columns: map[string]ddl.Column{
				"id":      {RawType: "BIGINT", NotNull: true},
				"user_id": {RawType: "BIGINT", NotNull: true},
			},
			PrimaryKey:  []string{"id"},
			ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}},
		},
	}
	out, err := RenderSchema(tables)
	if err != nil {
		t.Fatalf("RenderSchema error: %v", err)
	}
	for _, want := range []string{"datasource db {", "generator client {", "model User {", "model Post {"} {
		if !strings.Contains(out, want) {
			t.Errorf("schema missing %q", want)
		}
	}
}
