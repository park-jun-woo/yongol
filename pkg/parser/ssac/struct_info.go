//ff:type feature=ssac-parse type=model
//ff:what .ssac 파일에 선언된 Go struct 정보 타입
package ssac

// StructInfo는 .ssac 파일에 선언된 Go struct 정보다.
type StructInfo struct {
	Name   string        // "OnOrderCompletedMessage"
	Fields []StructField
}
