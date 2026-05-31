//ff:type feature=gen-nestjs type=model
//ff:what reverseRelation — 참조 모델이 Prisma 완전성을 위해 필요로 하는 역(has-many) 관계

package prisma

// reverseRelation tracks a reverse (has-many) relation that a referenced
// model needs for Prisma completeness.
type reverseRelation struct {
	FieldName string // plural snake_case (e.g. "posts")
	ModelName string // PascalCase referring model (e.g. "Post")
}
