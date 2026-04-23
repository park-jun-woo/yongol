//ff:type feature=migration type=model
//ff:what CreateTable — CREATE TABLE Operation
package migration

type CreateTable struct{ Table *Table }
