//ff:func feature=gen-ir type=util control=sequence
//ff:what convertAuth -- @auth 시퀀스 → AuthOp IR 변환 (Ownership 정보 이식)

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// convertAuth converts a @auth sequence to an IR Op. When the Fullstack
// context contains ParsedPolicies with @ownership annotations matching the
// sequence's Resource, AuthOp.Ownership is populated with the lookup metadata.
func convertAuth(seq ssac.Sequence, fs *yongol.Fullstack) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 403
	}
	op := AuthOp{
		Action:     seq.Action,
		Resource:   seq.Resource,
		Inputs:     convertInputsToFieldArgs(seq.Inputs),
		Message:    seq.Message,
		StatusCode: statusCode,
	}

	// Enrich with ownership info from Rego @ownership annotations.
	if fs != nil {
		for _, p := range fs.ParsedPolicies {
			for _, om := range p.Ownerships {
				if om.Resource == seq.Resource {
					info := OwnershipInfo{
						Table:       om.Table,
						OwnerColumn: om.Column,
					}
					// Resolve PK from DDL if available.
					if len(fs.DDLTables) > 0 {
						info.ResourcePK = findTablePK(fs, om.Table)
					}
					op.Ownership = &info
					break
				}
			}
			if op.Ownership != nil {
				break
			}
		}
	}

	return Op{Kind: OpAuth, Auth: &op}
}

// findTablePK returns the first primary key column name for the given DDL
// table name. Returns empty string if the table is not found or has no PK.
func findTablePK(fs *yongol.Fullstack, tableName string) string {
	for _, t := range fs.DDLTables {
		if strings.EqualFold(t.Name, tableName) {
			if len(t.PrimaryKey) > 0 {
				return t.PrimaryKey[0]
			}
			return ""
		}
	}
	return ""
}
