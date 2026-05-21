//ff:func feature=cli-init type=command control=sequence
//ff:what Run — materializes a minimal SSOT skeleton for `yongol init`

package cliinit

import (
	"fmt"
	"io"
	"strings"
)

// Run materializes the SSOT skeleton described in plans/cli/init/Phase001.
// Output is written to out (typically cmd.OutOrStdout()) and errOut (stderr)
// so the command is trivial to integration-test via cobra's SetOut/SetErr.
//
// When opts.FeaturesPath is set, features.yaml is parsed and SSOT stubs
// (OpenAPI paths, SSaC files, Rego rules, Hurl requests) are generated from
// the feature list. A specs/.yongol hash lock is also written.
func Run(out, errOut io.Writer, opts Options) error {
	if err := ValidateProjectID(opts.ProjectID); err != nil {
		return err
	}
	if opts.FeaturesPath == "" {
		return fmt.Errorf("features.yaml path is required")
	}

	// Parse features.yaml early to fail fast before any disk writes.
	ff, err := loadFeatures(opts.FeaturesPath)
	if err != nil {
		return err
	}

	// Description is optional when features.yaml is provided.
	if strings.TrimSpace(opts.Description) == "" {
		opts.Description = opts.ProjectID + " project"
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
	if err := createSkeletonDirs(opts.Dir); err != nil {
		return err
	}
	if err := writeSkeletonFiles(opts.Dir, data); err != nil {
		return err
	}

	// Features-driven stubs: overwrite the template-based openapi and rego
	// with feature-aware versions, and generate SSaC + Hurl stubs.
	if err := generateOpenAPIFromFeatures(opts.Dir, data, ff.Features); err != nil {
		return err
	}
	if err := generateSSaCFromFeatures(opts.Dir, ff.Features); err != nil {
		return err
	}
	if err := generateRegoFromFeatures(opts.Dir, ff.Features); err != nil {
		return err
	}
	if err := generateHurlFromFeatures(opts.Dir, ff.Features); err != nil {
		return err
	}
	if err := copyFeaturesYAML(opts.Dir, opts.FeaturesPath); err != nil {
		return err
	}
	if err := writeYongolHash(opts.Dir, opts.FeaturesPath); err != nil {
		return err
	}

	fmt.Fprintf(out, "yongol init: created %s (%d features)\n", opts.Dir, len(ff.Features))
	fmt.Fprintf(out, "  manifest.metadata.name = %s\n", data.ProjectIDNormalized)
	fmt.Fprintf(out, "  backend.module         = %s\n", data.Module)
	fmt.Fprintf(out, "\nNext: cd %s && yongol validate specs\n", opts.Dir)
	return nil
}
