//ff:type feature=validate type=model topic=ddl-structural
//ff:what positionalHit — $N 위치 파라미터 스캔 결과 (line + param 텍스트)

package ddl

// positionalHit captures a single $N positional parameter occurrence found
// while scanning a query body: its line number and the matched token.
type positionalHit struct {
	line  int
	param string
}
