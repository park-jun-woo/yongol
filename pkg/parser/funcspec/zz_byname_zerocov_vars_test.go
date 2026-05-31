package funcspec

const bnStructSrc = `package p
type FooRequest struct {
	Name string ` + "`json:\"name\"`" + `
	Age  int
}
type FooResponse struct {
	ID int
}
func g() {}
`
