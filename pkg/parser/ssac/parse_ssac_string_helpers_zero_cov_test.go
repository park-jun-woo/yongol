//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseBatch_ZeroCov — ssac 파서 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac

import (
	"testing"
)

func TestParseSsacStringHelpers_ZeroCov(t *testing.T) {
	// extractInputs / parseInputs
	if _, _, err := extractInputs("{ID: request.id}"); err != nil {
		t.Errorf("extractInputs: %v", err)
	}
	if _, err := parseInputs("{ID: request.id}"); err != nil {
		t.Errorf("parseInputs: %v", err)
	}
	// parseAnnotation dispatch
	if _, err := parseAnnotation("@get Course course = Course.FindByID({ID: request.id})"); err != nil {
		t.Errorf("parseAnnotation get: %v", err)
	}
	// parseCall
	if _, err := parseCall("auth.RefreshRotate({Token: request.token})"); err != nil {
		t.Errorf("parseCall: %v", err)
	}
	// parseEval
	if _, err := parseEval("schedule.IsExpired({At: request.at}) 400 \"past\""); err != nil {
		t.Errorf("parseEval: %v", err)
	}
	// parseCRUDWithResult
	var seq Sequence
	if err := parseCRUDWithResult("Course course = Course.FindByID({ID: request.id})", &seq); err != nil {
		t.Errorf("parseCRUDWithResult: %v", err)
	}
	// parseResponseLine / handleResponseLine
	if _, _, err := parseResponseLine("@response course"); err != nil {
		t.Errorf("parseResponseLine: %v", err)
	}
	handleResponseLine("@response course", nil, false)
	// parseResponseFields
	parseResponseFields([]string{"course: course", "name: instructor.Name"})
}
