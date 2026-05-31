//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestMapPGToPython(t *testing.T) {
	cases := []struct {
		raw     string
		notNull bool
		want    string
	}{
		{"BIGINT", true, "int"},
		{"BIGINT", false, "int | None"},
		{"VARCHAR(255)", true, "str"},
		{"TEXT[]", true, "list[str]"},
		{"INTEGER[]", false, "list[int] | None"},
		{"timestamptz", true, "datetime"},
	}
	for _, c := range cases {
		if got := mapPGToPython(c.raw, c.notNull); got != c.want {
			t.Errorf("mapPGToPython(%q,%v) = %q, want %q", c.raw, c.notNull, got, c.want)
		}
	}
}
