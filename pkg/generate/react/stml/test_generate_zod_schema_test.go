//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema — OpenAPI 제약조건에서 zod 스키마 코드 생성을 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaString(t *testing.T) {
	maxLen := 200
	fields := map[string]oapiparser.FieldConstraint{
		"title": {
			Type:      "string",
			Required:  true,
			MaxLength: &maxLen,
		},
		"trigger_event": {
			Type:     "string",
			Required: true,
		},
	}

	code := generateZodSchema("CreateWorkflow", fields)
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "title: z.string().min(1).max(200)")
	assertContains(t, code, "trigger_event: z.string().min(1)")
}

func TestGenerateZodSchemaEmpty(t *testing.T) {
	code := generateZodSchema("SomeOp", nil)
	if code != "" {
		t.Errorf("expected empty string for nil fields, got %q", code)
	}
	code = generateZodSchema("SomeOp", map[string]oapiparser.FieldConstraint{})
	if code != "" {
		t.Errorf("expected empty string for empty fields, got %q", code)
	}
}

func TestGenerateZodSchemaInteger(t *testing.T) {
	min := 1.0
	max := 100.0
	fields := map[string]oapiparser.FieldConstraint{
		"capacity": {
			Type:     "integer",
			Required: true,
			Minimum:  &min,
			Maximum:  &max,
		},
	}
	code := generateZodSchema("UpdateRoom", fields)
	assertContains(t, code, "capacity: z.number().int().min(1).max(100)")
}

func TestGenerateZodSchemaNumber(t *testing.T) {
	min := 0.0
	max := 99.99
	fields := map[string]oapiparser.FieldConstraint{
		"price": {
			Type:    "number",
			Minimum: &min,
			Maximum: &max,
		},
	}
	code := generateZodSchema("SetPrice", fields)
	assertContains(t, code, "price: z.number().min(0).max(99.99).optional()")
}

func TestGenerateZodSchemaBoolean(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"active": {
			Type:     "boolean",
			Required: true,
		},
	}
	code := generateZodSchema("Toggle", fields)
	assertContains(t, code, "active: z.boolean()")
	assertNotContains(t, code, ".optional()")
}

func TestGenerateZodSchemaEmail(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"email": {
			Type:     "string",
			Format:   "email",
			Required: true,
		},
	}
	code := generateZodSchema("Register", fields)
	assertContains(t, code, "email: z.string().email().min(1)")
}

func TestGenerateZodSchemaEnum(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"role": {
			Type:     "string",
			Enum:     []string{"admin", "member"},
			Required: true,
		},
	}
	code := generateZodSchema("Assign", fields)
	assertContains(t, code, `role: z.enum(["admin", "member"])`)
	assertNotContains(t, code, ".optional()")
}

func TestGenerateZodSchemaEnumOptional(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"role": {
			Type: "string",
			Enum: []string{"admin", "member"},
		},
	}
	code := generateZodSchema("Assign", fields)
	assertContains(t, code, `role: z.enum(["admin", "member"]).optional()`)
}

func TestGenerateZodSchemaPattern(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"phone": {
			Type:     "string",
			Required: true,
			Pattern:  `^\d{3}-\d{4}-\d{4}$`,
		},
	}
	code := generateZodSchema("UpdatePhone", fields)
	assertContains(t, code, `phone: z.string().min(1).regex(new RegExp("^\d{3}-\d{4}-\d{4}$"))`)
}

func TestGenerateZodSchemaOptionalField(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"bio": {
			Type: "string",
		},
	}
	code := generateZodSchema("UpdateProfile", fields)
	assertContains(t, code, "bio: z.string().optional()")
}

func TestGenerateZodSchemaMinLengthNoDoubleMin(t *testing.T) {
	// When required=true and minLength=1, only one .min(1) should appear.
	minLen := 1
	fields := map[string]oapiparser.FieldConstraint{
		"name": {
			Type:      "string",
			Required:  true,
			MinLength: &minLen,
		},
	}
	code := generateZodSchema("Create", fields)
	// Should have exactly one .min(1)
	count := strings.Count(code, ".min(1)")
	if count != 1 {
		t.Errorf("expected exactly 1 .min(1), got %d\n--- code ---\n%s", count, code)
	}
}

func TestGenerateZodSchemaMinLengthGreaterThan1(t *testing.T) {
	// When required=true and minLength=8, both .min(1) and .min(8) should appear.
	minLen := 8
	fields := map[string]oapiparser.FieldConstraint{
		"password": {
			Type:      "string",
			Required:  true,
			MinLength: &minLen,
		},
	}
	code := generateZodSchema("Login", fields)
	assertContains(t, code, "password: z.string().min(1).min(8)")
}

func TestGenerateZodSchemaDeterministicOrder(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"z_field": {Type: "string", Required: true},
		"a_field": {Type: "string", Required: true},
		"m_field": {Type: "string", Required: true},
	}
	code := generateZodSchema("Test", fields)
	aIdx := strings.Index(code, "a_field")
	mIdx := strings.Index(code, "m_field")
	zIdx := strings.Index(code, "z_field")
	if aIdx > mIdx || mIdx > zIdx {
		t.Errorf("fields not in alphabetical order\n--- code ---\n%s", code)
	}
}
