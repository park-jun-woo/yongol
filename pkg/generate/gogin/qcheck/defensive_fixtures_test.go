//ff:func feature=gen-gogin type=test control=sequence
//ff:what testDefensiveFixtures — DF-01/02/06 스캐너 테스트용 고정 소스 리터럴

package qcheck

// dfUnmarshalCheckedSrc matches the template emitted by
// generate_subscribe_method.go line 50 — a single-line IfStmt whose Init
// captures the Unmarshal error. Scanner must return zero findings.
const dfUnmarshalCheckedSrc = `package x

import "encoding/json"

func H(msg []byte) error {
	var m struct{ A int }
	if err := json.Unmarshal(msg, &m); err != nil { return err }
	return nil
}
`

// dfUnmarshalDiscardedSrc is a regression fixture that drops the error
// via `_ = json.Unmarshal(...)`. Scanner must flag DF-01.
const dfUnmarshalDiscardedSrc = `package x

import "encoding/json"

func H(msg []byte) {
	var m struct{ A int }
	_ = json.Unmarshal(msg, &m)
}
`

// dfScanCheckedSrc mirrors the sqlc-style two-line err := row.Scan(...)
// followed by an err-guard. Scanner must return zero findings.
const dfScanCheckedSrc = `package x

type row struct{}
func (r row) Scan(dest ...any) error { return nil }

func H(r row) error {
	var x int
	err := r.Scan(&x)
	if err != nil { return err }
	return nil
}
`

// dfScanDroppedSrc intentionally forgets the err check. Scanner must flag.
const dfScanDroppedSrc = `package x

type row struct{}
func (r row) Scan(dest ...any) error { return nil }

func H(r row) int {
	var x int
	_ = r.Scan(&x)
	return x
}
`

// dfDeferCloseOkSrc matches the boot/block_db_init.go template output.
// Scanner must return zero findings.
const dfDeferCloseOkSrc = `package x

type db struct{}
func (d *db) QueryContext(ctx any, q string) (*rows, error) { return nil, nil }
type rows struct{}
func (r *rows) Close() error { return nil }

func H(d *db, ctx any) error {
	rows, err := d.QueryContext(ctx, "select 1")
	if err != nil { return err }
	defer rows.Close()
	_ = rows
	return nil
}
`

// dfDeferCloseMissingSrc drops the defer. Scanner must flag DF-06.
const dfDeferCloseMissingSrc = `package x

type db struct{}
func (d *db) QueryContext(ctx any, q string) (*rows, error) { return nil, nil }
type rows struct{}
func (r *rows) Close() error { return nil }

func H(d *db, ctx any) error {
	rows, err := d.QueryContext(ctx, "select 1")
	if err != nil { return err }
	_ = rows
	return nil
}
`
