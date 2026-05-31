//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestPascalCase(t *testing.T) {
	cases := map[string]string{
		"users":       "Users",
		"order_items": "OrderItems",
		"":            "",
		"_leading":    "Leading",
		"a__b":        "AB",
		"single":      "Single",
		"trailing_":   "Trailing",
	}
	for in, want := range cases {
		if got := pascalCase(in); got != want {
			t.Errorf("pascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}
