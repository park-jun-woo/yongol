//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what overrideClassRe — @override 주석에서 class 값 추출 정규식
package stml_design

import (
	"regexp"
)

// overrideClassRe extracts the class value from @override comment text.
// Matches: @override class="..." or @override class='...'
var overrideClassRe = regexp.MustCompile("@override\\s+class=[\"']([^\"']+)[\"']")
