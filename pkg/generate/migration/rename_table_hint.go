//ff:type feature=migration type=model topic=migration-hints
//ff:what RenameTableHint — @rename 힌트에서 파생된 테이블 rename 매핑
package migration

// RenameTableHint maps an old table name to a new one.
type RenameTableHint struct{ From, To string }
