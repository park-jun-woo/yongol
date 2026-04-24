//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefEnum — dereferences a ~string-based enum pointer to string (nil→"")
//ff:checked llm=yongol-gen hash=cd3aab42
package service

func derefEnum[T ~string](p *T) string { if p != nil { return string(*p) }; return "" }
