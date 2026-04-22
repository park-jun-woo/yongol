//ff:type feature=migration type=model
//ff:what Operation 인터페이스 + 14개 구체 타입 (CREATE/DROP/ALTER + RENAME)
package migration

import (
	"fmt"
	"strings"
)

// SafetyLevel is the severity Phase004 attaches to each Operation.
type SafetyLevel int

const (
	SafetySafe SafetyLevel = iota
	SafetyWarning
	SafetyError
)

// Operation is one emitted migration step (one ALTER / CREATE / DROP).
type Operation interface {
	SQL() string
	Description() string
	// Destructive returns true for operations that can destroy data
	// (DROP TABLE/COLUMN, NOT NULL add without backfill, ...).
	Destructive() bool
	// SafetyLevel returns the classification used by check_safety.go.
	SafetyLevel() SafetyLevel
}

// ─────────────────────────────────────────────────────────────────────
// CREATE TABLE
// ─────────────────────────────────────────────────────────────────────

type CreateTable struct{ Table *Table }

func (op CreateTable) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "CREATE TABLE %s (", op.Table.Name)
	for i, c := range op.Table.Columns {
		if i == 0 {
			b.WriteString("\n    ")
		} else {
			b.WriteString(",\n    ")
		}
		b.WriteString(renderColumn(c))
	}
	if len(op.Table.PrimaryKey) > 0 {
		fmt.Fprintf(&b, ",\n    PRIMARY KEY (%s)", strings.Join(op.Table.PrimaryKey, ", "))
	}
	b.WriteString("\n);")
	return b.String()
}
func (op CreateTable) Description() string    { return "create table " + op.Table.Name }
func (op CreateTable) Destructive() bool      { return false }
func (op CreateTable) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// DROP TABLE
// ─────────────────────────────────────────────────────────────────────

type DropTable struct {
	Name             string
	AllowDestructive bool // set by check_safety from hints
}

func (op DropTable) SQL() string             { return fmt.Sprintf("DROP TABLE %s;", op.Name) }
func (op DropTable) Description() string     { return "drop table " + op.Name }
func (op DropTable) Destructive() bool       { return true }
func (op DropTable) SafetyLevel() SafetyLevel {
	if op.AllowDestructive {
		return SafetySafe
	}
	return SafetyWarning
}

// ─────────────────────────────────────────────────────────────────────
// ADD COLUMN
// ─────────────────────────────────────────────────────────────────────

type AddColumn struct {
	Table    string
	Column   *Column
	Backfill string // Phase004 — when set, prepend UPDATE before SET NOT NULL
}

func (op AddColumn) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "ALTER TABLE %s ADD COLUMN %s;",
		op.Table, renderColumn(op.Column))
	return b.String()
}
func (op AddColumn) Description() string    { return "add column " + op.Table + "." + op.Column.Name }
func (op AddColumn) Destructive() bool      { return false }
func (op AddColumn) SafetyLevel() SafetyLevel {
	// NOT NULL add without default or backfill is risky.
	if !op.Column.Nullable && op.Column.Default == "" && op.Backfill == "" {
		return SafetyError
	}
	return SafetySafe
}

// ─────────────────────────────────────────────────────────────────────
// DROP COLUMN
// ─────────────────────────────────────────────────────────────────────

type DropColumn struct {
	Table, Column    string
	AllowDestructive bool
}

func (op DropColumn) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", op.Table, op.Column)
}
func (op DropColumn) Description() string     { return "drop column " + op.Table + "." + op.Column }
func (op DropColumn) Destructive() bool       { return true }
func (op DropColumn) SafetyLevel() SafetyLevel {
	if op.AllowDestructive {
		return SafetySafe
	}
	return SafetyWarning
}

// ─────────────────────────────────────────────────────────────────────
// ALTER COLUMN …
// ─────────────────────────────────────────────────────────────────────

type AlterColumnType struct {
	Table, Column string
	From, To      CanonicalType
	Using         string // from @cast hint
}

func (op AlterColumnType) SQL() string {
	if op.Using != "" {
		return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s;",
			op.Table, op.Column, op.To.SQL(), op.Using)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s::%s;",
		op.Table, op.Column, op.To.SQL(), op.Column, strings.ToLower(op.To.Base))
}
func (op AlterColumnType) Description() string {
	return fmt.Sprintf("alter %s.%s type %s→%s", op.Table, op.Column, op.From.SQL(), op.To.SQL())
}
func (op AlterColumnType) Destructive() bool { return true }
func (op AlterColumnType) SafetyLevel() SafetyLevel {
	if op.Using != "" {
		return SafetySafe
	}
	return SafetyWarning
}

type AlterColumnNullable struct {
	Table, Column string
	From, To      bool
	Backfill      string // from @backfill hint
}

func (op AlterColumnNullable) SQL() string {
	verb := "SET NOT NULL"
	if op.To {
		verb = "DROP NOT NULL"
	}
	if op.To == false && op.Backfill != "" {
		return fmt.Sprintf("UPDATE %s SET %s = %s WHERE %s IS NULL;\nALTER TABLE %s ALTER COLUMN %s SET NOT NULL;",
			op.Table, op.Column, op.Backfill, op.Column, op.Table, op.Column)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s;", op.Table, op.Column, verb)
}
func (op AlterColumnNullable) Description() string {
	if op.To {
		return "drop NOT NULL on " + op.Table + "." + op.Column
	}
	return "set NOT NULL on " + op.Table + "." + op.Column
}
func (op AlterColumnNullable) Destructive() bool { return !op.To } // NOT NULL add is risky
func (op AlterColumnNullable) SafetyLevel() SafetyLevel {
	if !op.To && op.Backfill == "" {
		return SafetyError
	}
	return SafetySafe
}

type AlterColumnDefault struct {
	Table, Column string
	From, To      string
}

func (op AlterColumnDefault) SQL() string {
	if op.To == "" {
		return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s DROP DEFAULT;", op.Table, op.Column)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET DEFAULT %s;", op.Table, op.Column, op.To)
}
func (op AlterColumnDefault) Description() string {
	return "alter default " + op.Table + "." + op.Column + " → " + op.To
}
func (op AlterColumnDefault) Destructive() bool       { return false }
func (op AlterColumnDefault) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// INDEX
// ─────────────────────────────────────────────────────────────────────

type CreateIndex struct {
	Table string
	Index *Index
}

func (op CreateIndex) SQL() string {
	uniq := ""
	if op.Index.Unique {
		uniq = "UNIQUE "
	}
	where := ""
	if op.Index.Where != "" {
		where = " WHERE " + op.Index.Where
	}
	return fmt.Sprintf("CREATE %sINDEX %s ON %s (%s)%s;",
		uniq, op.Index.Name, op.Table,
		strings.Join(op.Index.Columns, ", "), where)
}
func (op CreateIndex) Description() string    { return "create index " + op.Index.Name }
func (op CreateIndex) Destructive() bool      { return false }
func (op CreateIndex) SafetyLevel() SafetyLevel { return SafetySafe }

type DropIndex struct{ Name string }

func (op DropIndex) SQL() string              { return fmt.Sprintf("DROP INDEX %s;", op.Name) }
func (op DropIndex) Description() string      { return "drop index " + op.Name }
func (op DropIndex) Destructive() bool        { return false }
func (op DropIndex) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// FK
// ─────────────────────────────────────────────────────────────────────

type AddForeignKey struct {
	Table string
	FK    *ForeignKey
}

func (op AddForeignKey) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)",
		op.Table, op.FK.Name,
		strings.Join(op.FK.Columns, ", "),
		op.FK.RefTable, strings.Join(op.FK.RefColumns, ", "))
	if op.FK.OnDelete != "" {
		fmt.Fprintf(&b, " ON DELETE %s", op.FK.OnDelete)
	}
	if op.FK.OnUpdate != "" {
		fmt.Fprintf(&b, " ON UPDATE %s", op.FK.OnUpdate)
	}
	b.WriteByte(';')
	return b.String()
}
func (op AddForeignKey) Description() string    { return "add FK " + op.FK.Name }
func (op AddForeignKey) Destructive() bool      { return false }
func (op AddForeignKey) SafetyLevel() SafetyLevel { return SafetySafe }

type DropForeignKey struct{ Table, Name string }

func (op DropForeignKey) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s;", op.Table, op.Name)
}
func (op DropForeignKey) Description() string     { return "drop FK " + op.Name }
func (op DropForeignKey) Destructive() bool       { return false }
func (op DropForeignKey) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// CHECK
// ─────────────────────────────────────────────────────────────────────

type AddCheck struct {
	Table string
	Check *CheckConstraint
}

func (op AddCheck) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s CHECK (%s);",
		op.Table, op.Check.Name, op.Check.Expression)
}
func (op AddCheck) Description() string     { return "add check " + op.Check.Name }
func (op AddCheck) Destructive() bool       { return false }
func (op AddCheck) SafetyLevel() SafetyLevel { return SafetySafe }

type DropCheck struct{ Table, Name string }

func (op DropCheck) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s;", op.Table, op.Name)
}
func (op DropCheck) Description() string     { return "drop check " + op.Name }
func (op DropCheck) Destructive() bool       { return false }
func (op DropCheck) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// RENAME (Phase004 힌트)
// ─────────────────────────────────────────────────────────────────────

type RenameColumn struct {
	Table, From, To string
}

func (op RenameColumn) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s;", op.Table, op.From, op.To)
}
func (op RenameColumn) Description() string     { return "rename column " + op.Table + "." + op.From + " → " + op.To }
func (op RenameColumn) Destructive() bool       { return false }
func (op RenameColumn) SafetyLevel() SafetyLevel { return SafetySafe }

type RenameTable struct {
	From, To string
}

func (op RenameTable) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", op.From, op.To)
}
func (op RenameTable) Description() string     { return "rename table " + op.From + " → " + op.To }
func (op RenameTable) Destructive() bool       { return false }
func (op RenameTable) SafetyLevel() SafetyLevel { return SafetySafe }

// ─────────────────────────────────────────────────────────────────────
// helpers
// ─────────────────────────────────────────────────────────────────────

// renderColumn emits a column clause for CREATE TABLE / ADD COLUMN.
func renderColumn(c *Column) string {
	b := strings.Builder{}
	b.WriteString(c.Name)
	b.WriteByte(' ')
	b.WriteString(c.Type.SQL())
	if !c.Nullable {
		b.WriteString(" NOT NULL")
	}
	if c.Default != "" {
		b.WriteString(" DEFAULT ")
		b.WriteString(c.Default)
	}
	return b.String()
}
