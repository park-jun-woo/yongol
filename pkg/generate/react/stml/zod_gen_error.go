//ff:type feature=stml-gen type=generator
//ff:what zodGenError — zod 스키마 생성 중 미지원 타입을 만났을 때의 명시적 실패 신호
package stml

import "fmt"

// zodGenError signals that the zod emitter encountered a type it cannot map
// safely. It is panicked from deep within the (string-returning) render chain
// and recovered at the GenerateWith boundary, where it is surfaced as a
// returned error — instead of silently downgrading the field to z.string().
type zodGenError struct {
	OperationID string
	Field       string
	Type        string
}

func (e *zodGenError) Error() string {
	op := e.OperationID
	if op == "" {
		op = "<unknown>"
	}
	field := e.Field
	if field == "" {
		field = "<unknown>"
	}
	return fmt.Sprintf("zod schema generation: unsupported field type %q for field %q in operation %q (no safe zod mapping)", e.Type, field, op)
}
