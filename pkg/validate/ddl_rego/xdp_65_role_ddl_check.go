//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XDP-65 — Rego role → DDL CHECK

package ddl_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdp65RoleDDLCheck validates XDP-65: every role value compared in Rego
// policies must be declared in the DDL role column CHECK(role IN (...))
// constraint. When the DDL has no role CHECK at all, the rule assumes no
// role model is defined and passes (same policy as the legacy
// bak/check_rego_role_ddl.go).
func xdp65RoleDDLCheck(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}

	// Collect allowed role values from any DDL table's CHECK on "role" column.
	allowed := make(map[string]bool)
	for _, t := range fs.DDLTables {
		c, ok := t.Columns["role"]
		if !ok {
			continue
		}
		for _, v := range c.CheckEnum {
			allowed[v] = true
		}
	}
	if len(allowed) == 0 {
		return nil
	}

	// Collect Rego role values with source context (first occurrence wins).
	type ctx struct {
		file string
		line int
	}
	regoRoles := make(map[string]ctx)
	for _, p := range fs.ParsedPolicies {
		for _, r := range p.Rules {
			if !r.UsesRole || r.RoleValue == "" {
				continue
			}
			if _, exists := regoRoles[r.RoleValue]; exists {
				continue
			}
			regoRoles[r.RoleValue] = ctx{file: p.File, line: r.SourceLine}
		}
	}

	var diags []diagnostic.Diagnostic
	for rv, c := range regoRoles {
		if allowed[rv] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  c.file,
			Line:  c.line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[XDP-65] Rego role %q is not defined in any DDL CHECK constraint",
				rv),
			Advice: fmt.Sprintf("Add '%s' to the role column CHECK IN constraint in the DDL user table", rv),
		})
	}
	return diags
}
