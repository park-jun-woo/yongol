//ff:type feature=stml-gen type=model
//ff:what 코드 생성 옵션을 설정하는 구조체
package stml

// GenerateOptions configures code generation behavior.
type GenerateOptions struct {
	APIImportPath string // import path for api module (default: "@/lib/api")
	UseClient     bool   // emit 'use client' directive (default: true)
}
