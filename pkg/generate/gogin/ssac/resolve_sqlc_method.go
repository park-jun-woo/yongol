//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolveSQLCMethod — SSaC Model.Method → sqlc Queries method name

package ssac

import "strings"

// resolveSQLCMethod converts "Workflow.FindByID" → "WorkflowFindByID".
// sqlc uses ModelPrefix + Method as query name.
func resolveSQLCMethod(model string) string {
	return strings.ReplaceAll(model, ".", "")
}
