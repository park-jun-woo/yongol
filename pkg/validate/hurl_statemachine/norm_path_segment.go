//ff:func feature=validate type=util control=selection topic=hurl-statemachine
//ff:what normPathSegment — 한 세그먼트를 원본 / `:param` 으로 정규화

package hurl_statemachine

import "regexp"

// normPathSegment returns ":param" when p is an OpenAPI-style
// placeholder (matched by openapiRe), a hurl `{{var}}`, or a pure
// integer literal. Otherwise returns p unchanged.
func normPathSegment(p string, openapiRe *regexp.Regexp) string {
	switch {
	case openapiRe != nil && openapiRe.MatchString(p):
		return ":param"
	case reHurlVarKey.MatchString(p):
		return ":param"
	case reHurlNumericK.MatchString(p):
		return ":param"
	default:
		return p
	}
}
