//ff:func feature=gen-fastapi type=generator control=iteration dimension=2
//ff:what writeSchemas — ServicePlan.BodyFields 기반 Pydantic BaseModel 스키마 파일 생성

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeSchemas generates Pydantic BaseModel schema files from ServicePlan
// BodyFields for each feature that has at least one POST/PUT/PATCH endpoint
// with a request body.
func writeSchemas(plansByFeature map[string][]*ir.ServicePlan, appDir string) error {
	schemasDir := filepath.Join(appDir, "schemas")
	if err := os.MkdirAll(schemasDir, 0o755); err != nil {
		return fmt.Errorf("mkdir schemas: %w", err)
	}

	for feature, plans := range plansByFeature {
		content := renderFeatureSchemas(plans)
		if content == "" {
			continue
		}
		path := filepath.Join(schemasDir, feature+".py")
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write schema %s: %w", feature, err)
		}
	}
	return nil
}

// renderFeatureSchemas produces Pydantic BaseModel classes for all plans in
// a feature that have request body fields. Returns empty string when no
// schemas are needed.
func renderFeatureSchemas(plans []*ir.ServicePlan) string {
	var b strings.Builder
	hasSchema := false

	for _, plan := range plans {
		method := strings.ToUpper(plan.HTTPMethod)
		hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
		if !hasBody {
			continue
		}
		if !hasSchema {
			b.WriteString("from pydantic import BaseModel\n")
			b.WriteString("from typing import Optional\n\n\n")
			hasSchema = true
		}
		renderOneSchema(&b, plan)
	}

	if !hasSchema {
		return ""
	}
	return b.String()
}

// renderOneSchema writes a single Pydantic BaseModel class for a plan's
// request body.
func renderOneSchema(b *strings.Builder, plan *ir.ServicePlan) {
	className := schemaPascalCase(plan.OperationID) + "Request"
	b.WriteString(fmt.Sprintf("class %s(BaseModel):\n", className))
	for _, field := range plan.BodyFields {
		pyType := schemaFormatToPython(field.Format)
		if field.Required {
			b.WriteString(fmt.Sprintf("    %s: %s\n", field.Name, pyType))
		} else {
			b.WriteString(fmt.Sprintf("    %s: Optional[%s] = None\n", field.Name, pyType))
		}
	}
	b.WriteString("\n\n")
}

// schemaFormatToPython maps an OpenAPI format string to a Python type.
func schemaFormatToPython(format string) string {
	switch format {
	case "email", "uuid", "uri", "url", "":
		return "str"
	case "date-time", "date":
		return "str"
	case "int32", "int64":
		return "int"
	case "float", "double":
		return "float"
	case "boolean":
		return "bool"
	default:
		return "str"
	}
}

// schemaPascalCase converts a camelCase or snake_case string to PascalCase.
func schemaPascalCase(s string) string {
	if s == "" {
		return ""
	}
	// If already PascalCase (starts with upper), return as-is.
	if s[0] >= 'A' && s[0] <= 'Z' {
		return s
	}
	// Convert first letter to uppercase for camelCase inputs.
	return strings.ToUpper(s[:1]) + s[1:]
}
