//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestMapPGToSA(t *testing.T) {
	cases := map[string]string{
		"BIGINT":       "Integer",
		"VARCHAR(255)": "String",
		"TEXT[]":       "ARRAY(Text)",
		"uuid":         "Uuid",
	}
	for in, want := range cases {
		if got := mapPGToSA(in); got != want {
			t.Errorf("mapPGToSA(%q) = %q, want %q", in, got, want)
		}
	}
}
