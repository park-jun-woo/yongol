//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefInt64 — dereferences *int64 to int64 (nil→0)
//ff:checked llm=yongol-gen hash=50743c23
package service

func derefInt64(p *int64) int64 { if p != nil { return *p }; return 0 }
