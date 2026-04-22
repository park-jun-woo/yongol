//ff:func feature=gen-gogin type=util control=sequence
//ff:what isCRUD — SSaC 시퀀스 타입이 CRUD (get/post/put/delete) 인지 판정

package ssac

// isCRUD reports whether seqType is one of the four CRUD verbs that map to
// sqlc-generated queries. Used by generateHTTPMethod to decide when an
// additional db import is needed.
func isCRUD(seqType string) bool {
	return seqType == "get" || seqType == "post" || seqType == "put" || seqType == "delete"
}
