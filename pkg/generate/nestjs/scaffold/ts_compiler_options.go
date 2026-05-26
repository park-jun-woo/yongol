//ff:type feature=gen-nestjs type=model
//ff:what TSCompilerOptions — TypeScript compilerOptions 구조체

package scaffold

// TSCompilerOptions holds the TypeScript compiler options subset needed
// for a NestJS project.
type TSCompilerOptions struct {
	Module                           string `json:"module"`
	Declaration                      bool   `json:"declaration"`
	RemoveComments                   bool   `json:"removeComments"`
	EmitDecoratorMetadata            bool   `json:"emitDecoratorMetadata"`
	ExperimentalDecorators           bool   `json:"experimentalDecorators"`
	AllowSyntheticDefaultImports     bool   `json:"allowSyntheticDefaultImports"`
	Target                           string `json:"target"`
	SourceMap                        bool   `json:"sourceMap"`
	OutDir                           string `json:"outDir"`
	BaseURL                          string `json:"baseUrl"`
	Incremental                      bool   `json:"incremental"`
	SkipLibCheck                     bool   `json:"skipLibCheck"`
	StrictNullChecks                 bool   `json:"strictNullChecks"`
	NoImplicitAny                    bool   `json:"noImplicitAny"`
	StrictBindCallApply              bool   `json:"strictBindCallApply"`
	ForceConsistentCasingInFileNames bool   `json:"forceConsistentCasingInFileNames"`
	NoFallthroughCasesInSwitch       bool   `json:"noFallthroughCasesInSwitch"`
}
