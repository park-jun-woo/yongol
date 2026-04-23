//ff:type feature=migration type=model topic=migration-hints
//ff:what RenameColumnHint — @rename 힌트에서 파생된 컬럼 rename 매핑
package migration

// RenameColumnHint maps an old column name to a new one inside a table.
type RenameColumnHint struct{ Table, From, To string }
