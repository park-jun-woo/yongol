//ff:func feature=migration type=test control=iteration dimension=1
//ff:what NormalizeType 30+ 케이스 — aliases·SERIAL·VARCHAR 길이·NUMERIC(p,s)·ARRAY
package migration

import "testing"

func TestNormalizeType(t *testing.T) {
	cases := []struct {
		in          string
		wantBase    string
		wantLen     int
		wantPrec    int
		wantScale   int
		wantArray   bool
		wantSerial  bool
	}{
		// Integer aliases
		{"int", "INTEGER", 0, 0, 0, false, false},
		{"int4", "INTEGER", 0, 0, 0, false, false},
		{"integer", "INTEGER", 0, 0, 0, false, false},
		{"INTEGER", "INTEGER", 0, 0, 0, false, false},
		{"bigint", "BIGINT", 0, 0, 0, false, false},
		{"int8", "BIGINT", 0, 0, 0, false, false},
		{"smallint", "SMALLINT", 0, 0, 0, false, false},
		{"int2", "SMALLINT", 0, 0, 0, false, false},

		// SERIAL → INTEGER/BIGINT with isSerial=true
		{"serial", "INTEGER", 0, 0, 0, false, true},
		{"BIGSERIAL", "BIGINT", 0, 0, 0, false, true},
		{"smallserial", "SMALLINT", 0, 0, 0, false, true},

		// Boolean aliases
		{"bool", "BOOLEAN", 0, 0, 0, false, false},
		{"boolean", "BOOLEAN", 0, 0, 0, false, false},

		// VARCHAR / character varying / char
		{"varchar(255)", "VARCHAR", 255, 0, 0, false, false},
		{"character varying(64)", "VARCHAR", 64, 0, 0, false, false},
		{"VARCHAR(1024)", "VARCHAR", 1024, 0, 0, false, false},
		{"char(10)", "CHAR", 10, 0, 0, false, false},
		{"character(5)", "CHAR", 5, 0, 0, false, false},

		// TEXT / UUID / JSONB
		{"text", "TEXT", 0, 0, 0, false, false},
		{"uuid", "UUID", 0, 0, 0, false, false},
		{"jsonb", "JSONB", 0, 0, 0, false, false},
		{"json", "JSON", 0, 0, 0, false, false},
		{"bytea", "BYTEA", 0, 0, 0, false, false},

		// Timestamp variants
		{"timestamp", "TIMESTAMP", 0, 0, 0, false, false},
		{"timestamp without time zone", "TIMESTAMP", 0, 0, 0, false, false},
		{"timestamptz", "TIMESTAMPTZ", 0, 0, 0, false, false},
		{"timestamp with time zone", "TIMESTAMPTZ", 0, 0, 0, false, false},
		{"date", "DATE", 0, 0, 0, false, false},

		// NUMERIC(p,s)
		{"numeric(10,2)", "NUMERIC", 0, 10, 2, false, false},
		{"decimal(5,0)", "NUMERIC", 0, 5, 0, false, false},
		{"numeric(7)", "NUMERIC", 0, 7, 0, false, false},

		// Array
		{"integer[]", "INTEGER", 0, 0, 0, true, false},
		{"text[]", "TEXT", 0, 0, 0, true, false},

		// Real / double precision
		{"real", "REAL", 0, 0, 0, false, false},
		{"double precision", "DOUBLE PRECISION", 0, 0, 0, false, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.in, func(t *testing.T) {
			got, isSerial := NormalizeType(c.in)
			if got.Base != c.wantBase {
				t.Errorf("Base: got %q, want %q", got.Base, c.wantBase)
			}
			if got.Length != c.wantLen {
				t.Errorf("Length: got %d, want %d", got.Length, c.wantLen)
			}
			if got.Precision != c.wantPrec {
				t.Errorf("Precision: got %d, want %d", got.Precision, c.wantPrec)
			}
			if got.Scale != c.wantScale {
				t.Errorf("Scale: got %d, want %d", got.Scale, c.wantScale)
			}
			if got.Array != c.wantArray {
				t.Errorf("Array: got %v, want %v", got.Array, c.wantArray)
			}
			if isSerial != c.wantSerial {
				t.Errorf("isSerial: got %v, want %v", isSerial, c.wantSerial)
			}
		})
	}
}
