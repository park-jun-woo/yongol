//ff:func feature=rule type=util control=sequence topic=catalog
//ff:what go:generate — repo 루트 rulebook.md 를 패키지 내 사본으로 복사 (//go:embed 는 상위 경로 불가)
package catalog

// go:embed 는 패키지 외부 파일을 참조할 수 없으므로, 릴리스 직전 또는 rulebook
// 편집 후 아래 명령으로 사본을 동기화한다:
//
//	go generate ./pkg/rule/catalog
//
// cmd/yongol/main.go 에서 MustLoad() 를 호출하는 시점에 누락이 드러나도록
// 설계돼 있어 동기화 누락은 기능 테스트 단계에서 즉시 감지된다.

//go:generate cp ../../../rulebook.md ./rulebook.md
