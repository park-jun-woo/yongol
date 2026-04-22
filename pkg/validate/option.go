//ff:type feature=validate type=model
//ff:what Option — Validate 동작 커스터마이징용 함수 타입
package validate

// Option customizes Validate behavior.
type Option func(*config)
