//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XQP-30 — @ownership 매핑은 대응 sqlc 쿼리 OwnerLookup<Resource> 가 존재해야 함

package query_rego

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqp30OwnerLookupQuery validates XQP-30: every Rego `@ownership` annotation
// demands a sqlc query whose name follows the `OwnerLookup<Resource>`
// convention declared in ssac/pkg/authz/interface.yaml.
//
// Naming contract:
//   @ownership workflow: workflows.org_id      →  OwnerLookupWorkflow
//   @ownership execution_log: execution_logs.org_id
//                                              →  OwnerLookupExecutionLog
//   @ownership lesson: courses.instructor_id via lessons.course_id
//                                              →  OwnerLookupLesson
//
// Resource names in `@ownership` are snake_case by convention; the sqlc
// query name is the PascalCase of the resource prefixed with
// `OwnerLookup`. The match is a simple set membership check against
// fs.SQLcQueries so renaming a query surfaces here as a missing mapping.
func xqp30OwnerLookupQuery(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}

	// Build a set of available sqlc query names for O(1) existence checks.
	have := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		have[q.Name] = true
	}

	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, p := range fs.ParsedPolicies {
		for _, om := range p.Ownerships {
			if om.Resource == "" {
				continue
			}
			want := ownerLookupName(om.Resource)
			key := p.File + "|" + want
			if seen[key] {
				continue
			}
			seen[key] = true
			if have[want] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  p.File,
				Line:  om.SourceLine,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelError,
				Message: fmt.Sprintf(
					"[XQP-30] @ownership %s — sqlc query %q not found; handler cannot load owner without it",
					om.Resource, want),
				Advice: buildAdvice(want, om.Table, om.Column, om.JoinTable, om.JoinFK),
			})
		}
	}
	return diags
}

// ownerLookupName builds the canonical sqlc query name for the supplied
// `@ownership` resource. Resource strings may arrive as snake_case
// (execution_log) or lowercase (workflow); either way the PascalCase form
// is produced by capitalising each segment.
func ownerLookupName(resource string) string {
	parts := strings.Split(resource, "_")
	var b strings.Builder
	b.WriteString("OwnerLookup")
	for _, part := range parts {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			b.WriteString(part[1:])
		}
	}
	return b.String()
}

// buildAdvice renders a ready-to-paste sqlc query stub. `via` annotations
// produce a JOIN form; direct mappings produce a single-table SELECT.
func buildAdvice(queryName, table, column, joinTable, joinFK string) string {
	if joinTable != "" && joinFK != "" {
		return fmt.Sprintf(
			"Add the following sqlc query to specs/db/queries/%s.sql:\n"+
				"-- name: %s :one\n"+
				"SELECT c.%s FROM %s c JOIN %s l ON l.%s = c.id WHERE l.id = @id;",
			joinTable, queryName, column, table, joinTable, joinFK)
	}
	return fmt.Sprintf(
		"Add the following sqlc query to specs/db/queries/%s.sql:\n"+
			"-- name: %s :one\n"+
			"SELECT %s FROM %s WHERE id = @id;",
		table, queryName, column, table)
}
