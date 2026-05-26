// Head token maps for PG type family classification.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package typemap

// integerHeads enumerates PG type tokens that map to the Integer family.
var integerHeads = map[string]bool{
	"BIGINT":      true,
	"BIGSERIAL":   true,
	"INTEGER":     true,
	"INT":         true,
	"INT4":        true,
	"INT8":        true,
	"INT2":        true,
	"SMALLINT":    true,
	"SERIAL":      true,
	"SMALLSERIAL": true,
}

// floatHeads enumerates PG type tokens that map to the Float family.
var floatHeads = map[string]bool{
	"REAL":   true,
	"FLOAT":  true,
	"FLOAT4": true,
	"FLOAT8": true,
}

// stringHeads enumerates PG type tokens that map to the String family.
var stringHeads = map[string]bool{
	"VARCHAR": true,
	"TEXT":    true,
	"CHAR":    true,
	"BPCHAR":  true,
}

// booleanHeads enumerates PG type tokens that map to the Boolean family.
var booleanHeads = map[string]bool{
	"BOOLEAN": true,
	"BOOL":    true,
}
