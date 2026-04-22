//ff:func feature=cli type=command control=sequence
//ff:what import — 외부 OpenAPI 문서에서 Go 클라이언트 모델을 생성한다.

package main

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/external"
	"github.com/spf13/cobra"
)

func importCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import <openapi-source> <output-dir>",
		Short: "Generate Go client model from an external OpenAPI document",
		Args:  usageArgs(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := external.Generate(args[0], args[1]); err != nil {
				return fmt.Errorf("import failed: %w", err)
			}
			return nil
		},
	}
}
