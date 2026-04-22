//ff:type feature=ssac-parse type=model
//ff:what 하나의 서비스 함수 선언 타입
package ssac

// ServiceFunc는 하나의 서비스 함수 선언이다.
type ServiceFunc struct {
	Name         string         // 함수명 (e.g. "GetCourse")
	FileName     string         // 원본 파일명
	Line         int            // 함수 선언 줄 번호 (1-based, 0 = 미상)
	Feature      string         // feature 폴더명 (e.g. "auth", 없으면 "")
	Sequences    []Sequence     // 시퀀스 목록
	Imports      []string       // Go import 경로
	Subscribe    *SubscribeInfo // nil이면 HTTP 트리거
	Param        *ParamInfo     // 함수 파라미터 (subscribe 함수용)
	Structs      []StructInfo   // .ssac 파일에 선언된 Go struct 목록
	NoPagination bool           // // @no-pagination 주석이 있으면 true — S-63 면제
}
