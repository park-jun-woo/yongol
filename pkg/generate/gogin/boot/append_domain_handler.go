//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what appendDomainHandler — 단일 도메인의 import + Group/publicOps/NewStrictHandler/RegisterHandlers 라인 추가

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// appendDomainHandler appends one domain's import and main.go lines to imports
// and lines, returning the grown slices. It mounts the domain at its
// route_prefix (r.Group), and when auth is active (authActive) builds a
// per-domain publicOps map from the domain's OpenAPI doc, wiring the strict
// middleware chosen by the domain's resolved auth_mode (Phase008 §3b): bearer →
// BearerAuthStrict<DomainTitle>, cookie/hybrid → CookieAuthStrict<DomainTitle>.
// The domain-suffixed func names avoid same-package redeclaration when two
// domains share a mode. The handler is registered on the route group with the
// shared srv (Decision B).
func appendDomainHandler(imports, lines []string, fs *yongol.Fullstack, name, modulePath string, authActive bool) ([]string, []string) {
	cfg := fs.Manifest.Domains[name]
	ident := domainIdent(name)
	title := domainTitle(name)
	pkg := "api_" + ident
	imports = append(imports, fmt.Sprintf(`"%s/internal/%s"`, modulePath, pkg))

	groupVar := ident + "Group"
	lines = append(lines, fmt.Sprintf("%s := r.Group(%q)", groupVar, cfg.RoutePrefix))

	// BUG-142 — mount this domain's request validator on its own route group so
	// payloads are validated against that domain's OpenAPI contract only. The
	// constructor returns (gin.HandlerFunc, error); bootstrap failures log and
	// os.Exit(1) (mirrors the single-site blockRequestValidator). middleware/fmt/
	// slog/os imports are emitted unconditionally for domain mode.
	imports = append(imports,
		fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
		`"fmt"`, `"log/slog"`, `"os"`,
	)
	validatorVar := ident + "Validator"
	lines = append(lines,
		fmt.Sprintf("%s, err := middleware.RequestValidator%s()", validatorVar, title),
		"if err != nil {",
		fmt.Sprintf("\tslog.Error(\"bootstrap failed\", \"stage\", \"request-validator-%s\", \"err\", err)", ident),
		fmt.Sprintf("\tfmt.Fprintf(os.Stderr, \"bootstrap failed: %%v\\n\", err)"),
		"\tos.Exit(1)",
		"}",
		fmt.Sprintf("%s.Use(%s)", groupVar, validatorVar),
	)

	var mwFuncs []string
	if authActive {
		opsVar := ident + "PublicOps"
		lines = append(lines, fmt.Sprintf("%s := map[string]bool{", opsVar))
		for _, opID := range collectPublicOps(fs.DomainView(name).OpenAPIDoc) {
			lines = append(lines, fmt.Sprintf("\t%q: true,", opID))
		}
		lines = append(lines, "}")
		title := domainTitle(name)
		switch domainWireMode(fs, name) {
		case "cookie", "hybrid":
			mwFuncs = append(mwFuncs, fmt.Sprintf("middleware.CookieAuthStrict%s(%s)", title, opsVar))
		default: // bearer
			mwFuncs = append(mwFuncs, fmt.Sprintf("middleware.BearerAuthStrict%s(%s)", title, opsVar))
		}
	}

	handlerVar := ident + "StrictHandler"
	lines = append(lines, fmt.Sprintf("%s := %s.NewStrictHandler(srv, []%s.StrictMiddlewareFunc{", handlerVar, pkg, pkg))
	for _, fn := range mwFuncs {
		lines = append(lines, "\t"+fn+",")
	}
	lines = append(lines, "})")
	lines = append(lines, fmt.Sprintf("%s.RegisterHandlers(%s, %s)", pkg, groupVar, handlerVar))
	return imports, lines
}
