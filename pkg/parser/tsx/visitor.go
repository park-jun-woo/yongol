//ff:type feature=tsx-parser type=model
//ff:what visitor — swc AST 순회 상태 (소스 바이트 + 라인 인덱스 + 누적 PageSpec)

package tsx

// visitor accumulates extraction results and maintains the byte-offset to
// (line, col) index derived from the source text.
type visitor struct {
	src        []byte
	lineOffset []int // byte offset of the start of each 1-based line
	page       *PageSpec
}
