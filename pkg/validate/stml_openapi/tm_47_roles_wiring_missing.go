//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-47 — data-roles 사용에 필요한 role 클레임 배선 부재 (role_field 미선언 / auth.claims 캡처 부재 / backend roles 빈 목록) (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm47RolesWiringMissing checks that a sitemap using data-roles has the
// full role-claim wiring (plans/stml/sitemap Phase005). The menu filter
// reads claims[<frontend.auth.role_field>], which an auth.claims.
// <role_field> data-capture fills from the login response, and TM-46
// validates the values against backend.auth.roles — so each missing link
// is an ERROR:
//   - frontend.auth.role_field is not declared (the filter would not know
//     which claim to read),
//   - no action captures auth.claims.<role_field> (the claim would never
//     be filled — every role-gated entry stays hidden forever), or
//   - backend.auth.roles is empty (no role vocabulary to validate
//     against; TM-46 stays silent then, deferring here).
//
// Without data-roles in the sitemap the rule is silent — the wiring is
// only required by its use.
func tm47RolesWiringMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	firstUse := firstRolesUse(fs.Sitemap)
	if firstUse == "" {
		return nil
	}

	var diags []diagnostic.Diagnostic
	roleField := ""
	if fs.Manifest != nil && fs.Manifest.Frontend.Auth != nil {
		roleField = fs.Manifest.Frontend.Auth.RoleField
	}
	if roleField == "" {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-47] sitemap uses data-roles (at %s) but manifest.yaml declares no frontend.auth.role_field — the menu filter would not know which claim holds the user's role", firstUse),
			Advice:  "Declare frontend.auth.role_field (e.g. role_field: role) naming the auth.claims.<name> entry the menu filter reads",
		})
	} else if !hasClaimsCapture(fs.STMLPages, roleField) {
		diags = append(diags, diagnostic.Diagnostic{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-47] sitemap uses data-roles (at %s) but no action captures \"auth.claims.%s\" — the role claim would never be filled and every role-gated menu entry would stay hidden", firstUse, roleField),
			Advice:  fmt.Sprintf("Add data-capture=\"<respField> -> auth.claims.%s\" to the login action (the field must exist in its 2xx response schema)", roleField),
		})
	}
	if len(backendAuthRoles(fs)) == 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-47] sitemap uses data-roles (at %s) but manifest backend.auth.roles is empty — there is no role vocabulary to validate the values against", firstUse),
			Advice:  "Declare the valid role names under backend.auth.roles (e.g. roles: [member, admin])",
		})
	}
	return diags
}
