//ff:type feature=validate type=model
//ff:what pageTokenRefs — 전체 STML 페이지에서 추출된 카테고리별 토큰 참조 집합
package stml_design

// pageTokenRefs holds all extracted custom token references from all pages.
type pageTokenRefs struct {
	Colors     []tokenRef // color-prefix tokens (bg-X, text-X, ...)
	Rounded    []tokenRef // rounded-X tokens
	Spacing    []tokenRef // spacing-prefix tokens (p-X, m-X, ...)
	Fonts      []tokenRef // font-X tokens
	Components []tokenRef // data-component references
}
