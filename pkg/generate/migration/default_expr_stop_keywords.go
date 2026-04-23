//ff:func feature=migration type=util control=sequence
//ff:what defaultExprStopKeywords — DEFAULT 표현식을 종료시키는 키워드 집합 반환
package migration

// defaultExprStopKeywords returns the set of uppercase SQL keywords
// that signal the end of a DEFAULT expression.
func defaultExprStopKeywords() map[string]bool {
	return map[string]bool{
		"NOT": true, "NULL": true, "UNIQUE": true, "PRIMARY": true,
		"REFERENCES": true, "CHECK": true, "DEFAULT": true, "CONSTRAINT": true,
	}
}
