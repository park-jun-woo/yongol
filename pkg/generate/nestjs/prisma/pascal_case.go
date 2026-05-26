//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what pascalCase — snake_case 테이블명 → PascalCase 모델명 변환

package prisma

import "strings"

// pascalCase converts a snake_case table name to PascalCase.
func pascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "")
}
