//ff:func feature=gen-nestjs type=util control=selection
//ff:what paginationOpt — pagination key/value → Prisma take/skip/cursor 옵션 문자열

package ssac

import "fmt"

// paginationOpt maps a single pagination key/value to its Prisma option string.
func paginationOpt(key, val string) string {
	switch key {
	case "per_page", "limit":
		return fmt.Sprintf("take: %s", val)
	case "page_offset", "offset":
		return fmt.Sprintf("skip: %s", val)
	case "cursor":
		return fmt.Sprintf("cursor: %s ? { id: %s } : undefined", val, val)
	default:
		return fmt.Sprintf("%s: %s", key, val)
	}
}
