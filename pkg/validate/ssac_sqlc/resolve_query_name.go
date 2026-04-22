//ff:func feature=validate type=util control=sequence topic=sqlc
//ff:what resolveQueryName — SSaC seq.Model → sqlc query name 변환

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// resolveQueryName converts SSaC "Workflow.FindByID" → "WorkflowFindByID" (sqlc query name).
func resolveQueryName(seq ssac.Sequence) string {
	return strings.ReplaceAll(seq.Model, ".", "")
}
