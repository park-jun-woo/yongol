//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefStr — dereferences *string to string (nil→"")
//ff:checked llm=yongol-gen hash=22a88a87
package service

func derefStr(p *string) string { if p != nil { return *p }; return "" }
