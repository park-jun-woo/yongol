//ff:type feature=gen-nestjs type=model
//ff:what TSConfig — TypeScript 컴파일러 설정 구조체

package scaffold

// TSConfig represents the TypeScript compiler configuration.
type TSConfig struct {
	CompilerOptions TSCompilerOptions `json:"compilerOptions"`
}
