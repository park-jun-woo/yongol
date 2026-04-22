//ff:func feature=cli type=command control=sequence
//ff:what versionCmd — yongol version 서브커맨드 반환
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print yongol version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "yongol %s\n", Version)
		},
	}
}
