//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what serverHelperSources — 포인터/deref 헬퍼 (fileName → 소스) 맵 반환

package ssac

// serverHelperSources returns (fileName → file body) for each helper. The
// content is pre-rendered with annotation blocks so the caller only owns
// the write step. Each entry is a single top-level func so filefunc F1
// (1 file 1 func) passes on the emitted service package.
func serverHelperSources() map[string]string {
	entries := helperSpecs()
	out := make(map[string]string, len(entries))
	for _, h := range entries {
		out[h.file] = renderHelperSource(h)
	}
	return out
}
