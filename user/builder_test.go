package user

// UserBuilder 是 Test Data Builder:預設值是一個合法的 User,
// 每個測試只改跟自己有關的欄位。
type UserBuilder struct{ u User }

func aUser() *UserBuilder {
	return &UserBuilder{u: User{
		Name: "test", Email: "t@example.com",
		Role: "member", Active: true,
	}}
}

func (b *UserBuilder) Role(r string) *UserBuilder {
	b.u.Role = r
	return b
}

func (b *UserBuilder) Inactive() *UserBuilder {
	b.u.Active = false
	return b
}

func (b *UserBuilder) Build() User { return b.u }
