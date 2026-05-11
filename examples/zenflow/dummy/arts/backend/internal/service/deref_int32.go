//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefInt32 — dereferences *int32 to int32 (nil→0)
//ff:checked llm=yongol-gen hash=d05f2044
package service

func derefInt32(p *int32) int32 { if p != nil { return *p }; return 0 }
