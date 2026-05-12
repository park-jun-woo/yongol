//ff:type feature=validate type=model
//ff:what overrideSet — 파일별 @override class 속성 값 집합
package stml_design

// overrideSet maps filename → set of class attribute values that are overridden.
type overrideSet map[string]map[string]bool
