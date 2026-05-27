//ff:func feature=gen-nestjs type=util control=sequence
//ff:what formatCallTarget — 외부 패키지 함수 → DI 서비스 메서드 참조 생성

package ssac

import "fmt"

// formatCallTarget formats a DI service method reference. When a package is
// specified, it becomes this.{pkg}Service.{lcMethod}(). Otherwise, it's a
// local method call.
func formatCallTarget(pkg, function string) string {
	method := lcFirst(function)
	if pkg == "" {
		return method
	}
	return fmt.Sprintf("this.%sService.%s", pkg, method)
}
