//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestInferDottedFieldType — var.field → SSaC.var→DDL.apifield/Struct 경로 타입 해석 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestInferDottedFieldType(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"SSaC.var.GetCourse.c":     "[]*models.Course",
			"DDL.apifield.Course.id":   "openapi_types.UUID",
			"Struct.Course.id":         "string",
			"Struct.Course.title":      "string",
			"SSaC.var.GetCourse.total": "int",
		},
	}

	// apifield preferred over Struct projection
	if got := inferDottedFieldType(g, "GetCourse", "c", "id"); got != "openapi_types.UUID" {
		t.Errorf("id: got %q, want openapi_types.UUID", got)
	}
	// falls back to Struct when no apifield key
	if got := inferDottedFieldType(g, "GetCourse", "c", "title"); got != "string" {
		t.Errorf("title: got %q, want string", got)
	}
	// unknown variable -> ""
	if got := inferDottedFieldType(g, "GetCourse", "unknown", "id"); got != "" {
		t.Errorf("unknown var: got %q, want empty", got)
	}
	// variable known but field has neither apifield nor Struct entry -> ""
	if got := inferDottedFieldType(g, "GetCourse", "total", "missing"); got != "" {
		t.Errorf("missing field: got %q, want empty", got)
	}
}
