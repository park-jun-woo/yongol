//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what ptrOf — wraps an arbitrary value T into *T (generic helper)
//ff:checked llm=yongol-gen hash=74a4ec29
package service

func ptrOf[T any](v T) *T { return &v }
