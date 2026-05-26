//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderVerifyPasswordOp — VerifyPasswordOp → SQLAlchemy lookup + bcrypt 검증 Python 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderVerifyPasswordOp writes a user lookup and password verification block.
func renderVerifyPasswordOp(b *strings.Builder, op *ir.VerifyPasswordOp, indent, sessionRef string) {
	if op == nil {
		return
	}
	model := pascalCase(op.Model)
	emailCol := snakeCase(op.EmailCol)
	b.WriteString(fmt.Sprintf("%sresult = await %s.execute(select(%s).where(%s.%s == %s))\n",
		indent, sessionRef, model, model, emailCol, op.EmailExpr))
	b.WriteString(fmt.Sprintf("%s%s = result.scalars().first()\n",
		indent, op.ResultVar))
	b.WriteString(fmt.Sprintf("%sif not %s:\n", indent, op.ResultVar))
	b.WriteString(fmt.Sprintf("%s    raise HTTPException(status_code=401, detail=\"%s\")\n",
		indent, op.Message))
	hashCol := snakeCase(op.HashCol)
	b.WriteString(fmt.Sprintf("%s# TODO: bcrypt.checkpw(%s, %s.%s)\n",
		indent, op.PasswordExpr, op.ResultVar, hashCol))
}
