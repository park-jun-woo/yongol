//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what missingBuiltinEntities — builtin backend 필수 DDL/쿼리 중 누락된 항목 나열

package manifest

// missingBuiltinEntities returns the list of required entities (DDL table,
// sqlc queries) that are absent from the user's artifacts for the given
// backend spec.
func missingBuiltinEntities(spec backendSpec, haveDDL, haveQuery map[string]bool) []string {
	var missing []string
	if spec.RequireDDL != "" && !haveDDL[spec.RequireDDL] {
		missing = append(missing, "DDL table "+spec.RequireDDL)
	}
	for _, qn := range spec.RequireQueries {
		if !haveQuery[qn] {
			missing = append(missing, "sqlc query "+qn)
		}
	}
	return missing
}
