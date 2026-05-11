//ff:func feature=stml-gen type=util control=sequence
//ff:what 기본 GenerateOptions 값을 반환한다
package stml

// DefaultOptions returns GenerateOptions with default values.
func DefaultOptions() GenerateOptions {
	return GenerateOptions{
		APIImportPath: "@/lib/api",
		UseClient:     true,
	}
}
