//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderNestCLI — nest-cli.json 생성

package scaffold

import (
	"encoding/json"
	"fmt"
)

// RenderNestCLI produces the nest-cli.json content.
func RenderNestCLI() (string, error) {
	cfg := NestCLIConfig{
		Schema:     "https://json.schemastore.org/nest-cli",
		Collection: "@nestjs/schematics",
		SourceRoot: "src",
		CompilerOptions: NestCompilerOpts{
			DeleteOutDir: true,
		},
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return "", fmt.Errorf("RenderNestCLI: %w", err)
	}
	return string(data) + "\n", nil
}
