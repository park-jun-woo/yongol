//ff:func feature=cli-init type=command control=sequence
//ff:what Run — materializes a minimal SSOT skeleton for `yongol init`

package cliinit

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// templateFiles bundles every template file shipped with the init command.
// Using embed.FS keeps templates inside the yongol binary so installations
// via `go install` work out-of-the-box without shipping a separate assets
// directory.
//
//go:embed templates/manifest.yaml.tmpl
//go:embed templates/openapi.yaml.tmpl
//go:embed templates/sqlc.yaml
//go:embed templates/authz.rego.tmpl
//go:embed templates/README.md.tmpl
//go:embed templates/gitignore
var templateFiles embed.FS

// Options is the parsed, validated input for Run. The CLI layer is expected
// to have already applied default values (e.g. Dir = "./<ProjectID>") and
// validated ProjectID via ValidateProjectID.
type Options struct {
	ProjectID   string // raw as given on the command line (PascalCase or snake_case)
	Description string // sentence suitable for manifest/openapi description
	Dir         string // destination directory; created if it does not exist
	Module      string // Go module path; if empty, DetectModule is invoked
	Force       bool   // allow writing into a non-empty directory
}

// Run materializes the SSOT skeleton described in plans/cli/init/Phase001.
// Output is written to out (typically cmd.OutOrStdout()) and errOut (stderr)
// so the command is trivial to integration-test via cobra's SetOut/SetErr.
func Run(out, errOut io.Writer, opts Options) error {
	if err := ValidateProjectID(opts.ProjectID); err != nil {
		return err
	}
	if strings.TrimSpace(opts.Description) == "" {
		return fmt.Errorf("description is required and must be non-empty")
	}
	if opts.Dir == "" {
		opts.Dir = "./" + opts.ProjectID
	}
	if opts.Module == "" {
		module, warning := DetectModule(opts.ProjectID)
		opts.Module = module
		if warning != "" {
			fmt.Fprintf(errOut, "yongol init: warning: %s\n", warning)
		}
	}
	if err := ensureEmptyDir(opts.Dir, opts.Force); err != nil {
		return err
	}

	data := templateData{
		ProjectID:           opts.ProjectID,
		ProjectIDNormalized: NormalizeProjectID(opts.ProjectID),
		Description:         opts.Description,
		Module:              opts.Module,
	}

	// Directory tree first so template writes cannot race the parents.
	for _, dir := range skeletonDirs() {
		abs := filepath.Join(opts.Dir, dir)
		if err := os.MkdirAll(abs, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", abs, err)
		}
	}

	for _, f := range skeletonFiles() {
		dest := filepath.Join(opts.Dir, f.destRel)
		var contents []byte
		if f.rendered {
			rendered, err := renderTemplate(f.srcEmbed, data)
			if err != nil {
				return fmt.Errorf("render %s: %w", f.srcEmbed, err)
			}
			contents = rendered
		} else {
			raw, err := templateFiles.ReadFile(f.srcEmbed)
			if err != nil {
				return fmt.Errorf("read embedded %s: %w", f.srcEmbed, err)
			}
			contents = raw
		}
		if err := os.WriteFile(dest, contents, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", dest, err)
		}
	}

	fmt.Fprintf(out, "yongol init: created %s\n", opts.Dir)
	fmt.Fprintf(out, "  manifest.metadata.name = %s\n", data.ProjectIDNormalized)
	fmt.Fprintf(out, "  backend.module         = %s\n", data.Module)
	fmt.Fprintf(out, "\nNext: cd %s && yongol validate specs\n", opts.Dir)
	return nil
}

// skeletonDirs returns the directory tree that must exist regardless of
// whether any template file lands in it. Ordering is from shallow to deep so
// os.MkdirAll never has to backfill intermediate nodes in later entries.
func skeletonDirs() []string {
	return []string{
		"specs",
		"specs/api",
		"specs/db",
		"specs/db/queries",
		"specs/service",
		"specs/states",
		"specs/policy",
		"specs/frontend",
		"specs/frontend/pages",
		"specs/frontend/components",
		"specs/tests",
	}
}

// skeletonFile describes one file to materialize under the target directory.
// `rendered=true` runs the file through text/template with templateData;
// `rendered=false` copies the embedded bytes verbatim (used for files that
// contain no placeholders, like sqlc.yaml).
type skeletonFile struct {
	srcEmbed string // path within the embed.FS
	destRel  string // path relative to opts.Dir
	rendered bool
}

func skeletonFiles() []skeletonFile {
	return []skeletonFile{
		{"templates/manifest.yaml.tmpl", "specs/manifest.yaml", true},
		{"templates/openapi.yaml.tmpl", "specs/api/openapi.yaml", true},
		{"templates/sqlc.yaml", "specs/db/sqlc.yaml", false},
		{"templates/authz.rego.tmpl", "specs/policy/authz.rego", true},
		{"templates/README.md.tmpl", "README.md", true},
		// gitignore is stored without a leading dot so `go:embed` does not
		// silently skip it (embed ignores files whose name begins with "." or
		// "_"). We rename on write.
		{"templates/gitignore", ".gitignore", false},
	}
}

// templateData is the context passed to text/template when rendering
// skeleton files. Keep fields exported so template authors can reference
// them via `{{.ProjectID}}` and friends.
type templateData struct {
	ProjectID           string
	ProjectIDNormalized string
	Description         string
	Module              string
}

// renderTemplate loads a template from the embed.FS and executes it against
// the supplied data. A fresh *template.Template is built per call to keep the
// function easy to reason about; the cost is negligible (N templates, each
// parsed once per `yongol init` invocation).
func renderTemplate(srcEmbed string, data templateData) ([]byte, error) {
	raw, err := templateFiles.ReadFile(srcEmbed)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(filepath.Base(srcEmbed)).Parse(string(raw))
	if err != nil {
		return nil, err
	}
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

// ensureEmptyDir refuses to touch a directory that already contains files
// unless --force is set. A dedicated helper keeps Run() readable and makes
// the semantic easy to unit-test.
func ensureEmptyDir(dir string, force bool) error {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", dir)
	}
	if force {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", dir, err)
	}
	if len(entries) > 0 {
		return fmt.Errorf("%s is not empty (use -f/--force to override)", dir)
	}
	return nil
}
