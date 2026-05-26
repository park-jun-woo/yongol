//ff:type feature=gen-nestjs type=model
//ff:what NestCompilerOpts — NestJS CLI 컴파일러 옵션 구조체

package scaffold

// NestCompilerOpts holds NestJS CLI compiler options.
type NestCompilerOpts struct {
	DeleteOutDir bool `json:"deleteOutDir"`
}
