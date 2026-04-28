// templateFiles bundles every template file shipped with the init command.
// Using embed.FS keeps templates inside the yongol binary so installations
// via `go install` work out-of-the-box without shipping a separate assets
// directory.

package cliinit

import "embed"

//go:embed templates/manifest.yaml.tmpl
//go:embed templates/openapi.yaml.tmpl
//go:embed templates/sqlc.yaml
//go:embed templates/authz.rego.tmpl
//go:embed templates/README.md.tmpl
//go:embed templates/gitignore
var templateFiles embed.FS
