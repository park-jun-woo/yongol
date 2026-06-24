//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.dbToAPIConvert — db 필드를 api 타입으로 변환하는 표현식 생성

package ssac

import "fmt"

// dbToAPIConvert generates a conversion expression from db field to api field.
// e.g. db.Workflow{ID: 1, Title: "x"} → api.Workflow{Id: &row.ID, Title: &row.Title}
func (g *methodGen) dbToAPIConvert(varName, refType string) string {
	return fmt.Sprintf("convert%s%s(%s)", g.DomainTitle, refType, varName)
}
