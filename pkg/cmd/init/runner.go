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
	if err := createSkeletonDirs(opts.Dir); err != nil {
		return err
	}
	if err := writeSkeletonFiles(opts.Dir, data); err != nil {
		return err
	}
	fmt.Fprintf(out, "yongol init: created %s\n", opts.Dir)
	fmt.Fprintf(out, "  manifest.metadata.name = %s\n", data.ProjectIDNormalized)
	fmt.Fprintf(out, "  backend.module         = %s\n", data.Module)
	fmt.Fprintf(out, "\nNext: cd %s && yongol validate specs\n", opts.Dir)
	return nil
}
