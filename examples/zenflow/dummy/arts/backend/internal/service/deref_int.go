//ff:func feature=service type=util control=sequence topic=pointer-helper
//ff:what derefInt — dereferences *int to int32 (nil→0)
//ff:checked llm=yongol-gen hash=3df37e5b
package service

func derefInt(p *int) int32 { if p != nil { return int32(*p) }; return 0 }
