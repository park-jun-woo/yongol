//ff:type feature=generate type=model
//ff:what GenerateOption — Generate 시그니처를 깨지 않고 옵션을 주입하는 functional option 타입
package generate

// GenerateOption tunes Generate without breaking the existing signature.
type GenerateOption func(*generateConfig)
