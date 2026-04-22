//ff:type feature=validate type=model topic=ddl-structural
//ff:what positionalHit — $N positional parameter scan result (line number + matched token)

package ddl

// positionalHit captures a single $N positional parameter occurrence found
// while scanning a query body: its line number and the matched token.
type positionalHit struct {
	line  int
	param string
}
