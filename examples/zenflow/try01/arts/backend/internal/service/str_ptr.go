//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what strPtr — wraps a string literal into *string
//ff:checked llm=yongol-gen hash=fef4f621
package service

func strPtr(s string) *string { return &s }
