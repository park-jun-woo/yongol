//ff:type feature=gen-nestjs type=model
//ff:what externalPackage — @call/@eval 외부 패키지와 메서드 목록

package nestjs

// externalPackage groups the methods referenced from a single external package.
type externalPackage struct {
	Name    string
	Methods []string
}
