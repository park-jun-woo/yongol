//ff:func feature=rule type=accessor control=sequence topic=catalog
//ff:what Source — 내장된 rulebook 원본 바이트 노출 (테스트에서 직접 재파싱용)
package catalog

// Source exposes the raw embedded rulebook bytes. Tests use this to
// re-parse the canonical source directly.
func Source() []byte {
	return rulebookSource
}
