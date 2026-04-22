//ff:func feature=gen-filefunc type=util control=sequence
//ff:what hasInfraDescription — 주어진 패키지명이 고정 infra 설명 테이블에 있는지 확인
package filefunc

// hasInfraDescription reports whether the given feature name has a fixed
// infrastructure description baked into infraDescriptions.
func hasInfraDescription(name string) bool {
	_, ok := infraDescriptions[name]
	return ok
}
