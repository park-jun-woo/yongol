//ff:type feature=gen-nestjs type=model
//ff:what NestCLIConfig — nest-cli.json 설정 구조체

package scaffold

// NestCLIConfig represents the nest-cli.json configuration.
type NestCLIConfig struct {
	Schema          string           `json:"$schema"`
	Collection      string           `json:"collection"`
	SourceRoot      string           `json:"sourceRoot"`
	CompilerOptions NestCompilerOpts `json:"compilerOptions"`
}
