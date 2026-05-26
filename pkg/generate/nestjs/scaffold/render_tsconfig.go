//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderTSConfig — NestJS tsconfig.json 생성

package scaffold

import (
	"encoding/json"
	"fmt"
)

// RenderTSConfig produces the tsconfig.json content for a NestJS project.
func RenderTSConfig() (string, error) {
	cfg := TSConfig{
		CompilerOptions: TSCompilerOptions{
			Module:                           "commonjs",
			Declaration:                      true,
			RemoveComments:                   true,
			EmitDecoratorMetadata:            true,
			ExperimentalDecorators:           true,
			AllowSyntheticDefaultImports:     true,
			Target:                           "ES2021",
			SourceMap:                        true,
			OutDir:                           "./dist",
			BaseURL:                          "./",
			Incremental:                      true,
			SkipLibCheck:                     true,
			StrictNullChecks:                 true,
			NoImplicitAny:                    true,
			StrictBindCallApply:              true,
			ForceConsistentCasingInFileNames: true,
			NoFallthroughCasesInSwitch:       true,
		},
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return "", fmt.Errorf("RenderTSConfig: %w", err)
	}
	return string(data) + "\n", nil
}
