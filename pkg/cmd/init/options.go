//ff:type feature=cli-init type=model
//ff:what Options — parsed, validated input for Run

package cliinit

// Options is the parsed, validated input for Run. The CLI layer is expected
// to have already applied default values (e.g. Dir = "./<ProjectID>") and
// validated ProjectID via ValidateProjectID.
type Options struct {
	ProjectID    string // raw as given on the command line (PascalCase or snake_case)
	FeaturesPath string // path to features.yaml (required); parsed to generate SSOT stubs
	Description  string // sentence suitable for manifest/openapi description (optional when features.yaml has entries)
	Dir          string // destination directory; created if it does not exist
	Module       string // Go module path; if empty, DetectModule is invoked
	Force        bool   // allow writing into a non-empty directory
}
