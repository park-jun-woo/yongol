//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFuncs — 활성 블록의 Funcs 를 순서대로 수집

package boot

// collectFuncs gathers top-level function declarations from all blocks in
// order. Unlike imports, Funcs are not deduplicated — each entry is a named
// declaration whose uniqueness is the block author's responsibility.
func collectFuncs(blocks []MainBlock) []string {
	var out []string
	for _, b := range blocks {
		out = append(out, b.Funcs...)
	}
	return out
}
