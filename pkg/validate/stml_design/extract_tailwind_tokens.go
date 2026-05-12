//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what tailwindPrefixes — Tailwind 커스텀 토큰 추출에 사용되는 prefix 목록 및 정규식 정의
package stml_design

import (
	"regexp"
)

// colorPrefixes lists Tailwind utility prefixes that reference color tokens.
var colorPrefixes = []string{
	"bg-", "text-", "border-", "ring-", "shadow-", "outline-",
	"accent-", "fill-", "stroke-", "divide-",
	"from-", "via-", "to-",
	"decoration-", "placeholder-",
}

// spacingPrefixes lists Tailwind utility prefixes that reference spacing tokens.
var spacingPrefixes = []string{
	"p-", "px-", "py-", "pt-", "pr-", "pb-", "pl-", "ps-", "pe-",
	"m-", "mx-", "my-", "mt-", "mr-", "mb-", "ml-", "ms-", "me-",
	"gap-", "gap-x-", "gap-y-",
	"space-x-", "space-y-",
	"w-", "h-", "min-w-", "min-h-", "max-w-", "max-h-",
	"inset-", "top-", "right-", "bottom-", "left-",
	"basis-", "size-",
}

// tailwindPaletteRe matches standard Tailwind color palette names (e.g. "gray-500", "red-200").
var tailwindPaletteRe = regexp.MustCompile(`^[a-z]+-\d+$`)
