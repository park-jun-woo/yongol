//ff:type feature=migration type=model
//ff:what Mode — Generate 가 반환하는 실행 브랜치 표기 (initial/incremental/noop)
package migration

// Mode indicates which branch the generator took.
type Mode string

const (
	ModeInitial     Mode = "initial"
	ModeIncremental Mode = "incremental"
	ModeNoop        Mode = "noop"
)
