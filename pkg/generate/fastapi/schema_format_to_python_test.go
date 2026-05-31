//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteSchemas — feature 별 Pydantic 스키마 파일 기록 검증
package fastapi

import (
	"testing"
)

func TestSchemaFormatToPython(t *testing.T) {
	cases := map[string]string{
		"email":     "str",
		"uuid":      "str",
		"uri":       "str",
		"url":       "str",
		"":          "str",
		"date-time": "str",
		"date":      "str",
		"int32":     "int",
		"int64":     "int",
		"float":     "float",
		"double":    "float",
		"boolean":   "bool",
		"unknown":   "str",
	}
	for format, want := range cases {
		if got := schemaFormatToPython(format); got != want {
			t.Errorf("schemaFormatToPython(%q) = %q, want %q", format, got, want)
		}
	}
}
