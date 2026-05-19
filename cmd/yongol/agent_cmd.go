//ff:func feature=cli type=command control=sequence
//ff:what agentCmd — yongol agent 서브커맨드 (validate 0 errors까지 SSOT 자동 수정)

package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/cmd/agent"
)

func agentCmd() *cobra.Command {
	var modelFlag string
	var maxRoundsFlag int
	var docsFlag string

	cmd := &cobra.Command{
		Use:           "agent <specs-dir>",
		Short:         "Auto-fix SSOT files until validate reports 0 errors",
		Args:          usageArgs(cobra.ExactArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			backend, model, err := parseModelFlag(modelFlag)
			if err != nil {
				return &usageError{err: err}
			}

			cfg := agent.Config{
				SpecsDir:  args[0],
				DocsDir:   docsFlag,
				Backend:   backend,
				Model:     model,
				MaxRounds: maxRoundsFlag,
			}

			return agent.Run(cmd.OutOrStdout(), cfg)
		},
	}

	cmd.Flags().StringVar(&modelFlag, "model", "ollama:gemma4:e4b",
		"LLM backend and model (format: ollama:<name>, xai:<name>, gemini:<name>)")
	cmd.Flags().IntVar(&maxRoundsFlag, "max-rounds", 20,
		"maximum validate-fix rounds")
	cmd.Flags().StringVar(&docsFlag, "docs", "",
		"path to yongol docs/ directory (auto-detected if empty)")

	return cmd
}

// parseModelFlag splits "backend:model" into backend and model name.
// For ollama models with colons in the name (e.g. "ollama:gemma4:e4b"),
// the backend is the first segment and the model is everything after.
func parseModelFlag(flag string) (backend, model string, err error) {
	idx := strings.Index(flag, ":")
	if idx < 0 {
		return "", "", fmt.Errorf("invalid --model %q: expected format backend:model (e.g. ollama:gemma4:e4b)", flag)
	}
	backend = flag[:idx]
	model = flag[idx+1:]
	switch backend {
	case "ollama", "xai", "gemini":
		// ok
	default:
		return "", "", fmt.Errorf("invalid --model backend %q: supported backends: ollama, xai, gemini", backend)
	}
	if model == "" {
		return "", "", fmt.Errorf("invalid --model %q: model name is empty", flag)
	}
	return backend, model, nil
}
