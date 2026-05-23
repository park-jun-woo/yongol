//ff:func feature=cli type=command control=sequence
//ff:what agentCmd — yongol agent 서브커맨드 (validate 0 errors까지 SSOT 자동 수정)

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/cmd/agent"
)

func agentCmd() *cobra.Command {
	var modelFlag string
	var maxRoundsFlag int

	cmd := &cobra.Command{
		Use:           "agent <specs-dir>",
		Short:         "Auto-fix SSOT files until validate reports 0 errors",
		Args:          usageArgs(cobra.ExactArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.ErrOrStderr(), "⚠ yongol agent는 실험 버전입니다. 생성된 코드가 불완전할 수 있습니다.")

			backend, model, err := parseModelFlag(modelFlag)
			if err != nil {
				return &usageError{err: err}
			}

			cfg := agent.Config{
				SpecsDir:  args[0],
				Backend:   backend,
				Model:     model,
				MaxRounds: maxRoundsFlag,
			}

			return agent.Run(cmd.OutOrStdout(), cfg)
		},
	}

	cmd.Flags().StringVar(&modelFlag, "model", "ollama:gemma4:e4b",
		"LLM backend and model (format: ollama:<name>, xai:<name>, gemini:<name>)")
	cmd.Flags().IntVar(&maxRoundsFlag, "max-rounds", 10,
		"maximum validate-fix rounds")

	return cmd
}
