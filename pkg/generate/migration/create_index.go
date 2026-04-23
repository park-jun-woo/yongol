//ff:type feature=migration type=model
//ff:what CreateIndex — CREATE INDEX Operation
package migration

type CreateIndex struct {
	Table string
	Index *Index
}
