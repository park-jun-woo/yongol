//ff:func feature=gen-gogin type=accessor control=sequence
//ff:what methodGen.queryVar — sqlc queries 변수명 반환 (tx 유무에 따라)

package ssac

// queryVar returns the sqlc queries variable name.
func (g *methodGen) queryVar() string {
	if g.UseTx {
		return "qtx"
	}
	return "server.Queries"
}
