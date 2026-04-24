//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefBool — dereferences *bool to bool (nil→false)
//ff:checked llm=yongol-gen hash=7159b9a7
package service

func derefBool(p *bool) bool { if p != nil { return *p }; return false }
