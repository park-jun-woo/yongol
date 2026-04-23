//ff:type feature=report type=model topic=sarif
//ff:what Message — 사람이 읽는 진단 텍스트 래퍼
package sarif

// Message carries the human-readable finding text.
type Message struct {
	Text string `json:"text"`
}
