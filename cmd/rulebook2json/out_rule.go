//ff:type feature=rule type=model
//ff:what outRule — quest seed 가 소비하는 rules.json 한 행의 JSON 형태 (lowercase 안정 키: id/level/desc/source/section)
package main

// outRule is the JSON shape consumed by the quest seed: lowercase, stable keys.
type outRule struct {
	ID      string `json:"id"`
	Level   string `json:"level"`   // "ERROR" | "WARNING"
	Desc    string `json:"desc"`    // rulebook Description column
	Source  string `json:"source"`  // Go source path of the current (legacy) rule
	Section string `json:"section"` // H2 section title, e.g. "A. SSaC Internal"
}
