//ff:type feature=gen-gogin type=model
//ff:what columnAdapter — ddl.Column 을 typemap.ColumnMeta 인터페이스에 맞추는 어댑터

package types

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// columnAdapter wraps ddl.Column to satisfy typemap.ColumnMeta.
type columnAdapter struct{ col ddl.Column }

func (a columnAdapter) RawType() string     { return a.col.RawType }
func (a columnAdapter) CheckEnum() []string { return a.col.CheckEnum }
