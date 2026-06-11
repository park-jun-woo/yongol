//ff:func feature=validate type=rule control=sequence topic=openapi-structural
//ff:what xoe02Detail — XOE-02 진단의 라인·메시지를 표시 필드 부재/비-string 경우로 나눠 산출한다

package openapi

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// xoe02Detail resolves the diagnostic line and message for schema name. When
// no "error"/"message" property exists at all the handler can only ever fall
// back to String(err); when the property exists but is non-string the
// `typeof msg === 'string'` guard still drops to String(err).
func xoe02Detail(name string, hasError, hasMessage bool, lines *oapiparser.LineIndex) (int, string) {
	if !hasError && !hasMessage {
		msg := fmt.Sprintf(
			"[XOE-02] ErrorResponse schema %q에 string 속성 `error`(차선 `message`)가 없어, 생성 프론트엔드가 모든 액션 실패를 String(err) 폴백(\"[object Object]\")으로만 표시합니다. string `error` 속성을 추가하세요.",
			name)
		return lines.SchemaLine(name), msg
	}

	field := "error"
	if !hasError && hasMessage {
		field = "message"
	}
	line := lines.SchemaPropertyLine(name, field)
	if line == 0 {
		line = lines.SchemaLine(name)
	}
	msg := fmt.Sprintf(
		"[XOE-02] ErrorResponse schema %q의 표시 필드 `%s`가 string 타입이 아니어서 생성 프론트엔드가 `typeof msg === 'string'` 가드에 걸려 String(err) 폴백으로 표시합니다. `%s`를 type: string으로 선언하세요.",
		name, field, field)
	return line, msg
}
