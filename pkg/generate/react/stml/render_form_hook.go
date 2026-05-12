//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useForm 훅 호출 코드를 생성한다 (zod 스키마가 있으면 zodResolver 적용)
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFormHook generates a useForm hook call.
// When constraints are available for the action's operationId, a zod schema is
// generated and zodResolver is applied.
func renderFormHook(a stmlparser.ActionBlock, constraints map[string]map[string]oapiparser.FieldConstraint) string {
	formName := toLowerFirst(a.OperationID) + "Form"

	fields := lookupConstraints(a.OperationID, constraints)
	if len(fields) == 0 {
		return fmt.Sprintf(`const %s = useForm()`, formName)
	}

	schemaCode := generateZodSchema(a.OperationID, fields)
	schemaName := toLowerFirst(a.OperationID) + "Schema"
	return fmt.Sprintf("%s\n  const %s = useForm({\n    resolver: zodResolver(%s),\n  })", schemaCode, formName, schemaName)
}

// lookupConstraints retrieves the field constraints for the given operationId.
// Returns nil when no constraints exist.
func lookupConstraints(operationID string, constraints map[string]map[string]oapiparser.FieldConstraint) map[string]oapiparser.FieldConstraint {
	if constraints == nil {
		return nil
	}
	fields, ok := constraints[operationID]
	if !ok || len(fields) == 0 {
		return nil
	}
	return fields
}
