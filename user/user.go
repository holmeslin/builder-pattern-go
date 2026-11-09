// Package user 是情境 3「測試資料」要測的程式。
// Test Data Builder 本身放在 builder_test.go:它只給測試用,不進正式程式碼。
package user

type User struct {
	Name, Email, Role string
	Active            bool
}

// CanDelete:只有啟用中的 admin 可以刪除。
func CanDelete(u User) bool { return u.Active && u.Role == "admin" }

// CanLogin:停用的帳號不能登入。
func CanLogin(u User) bool { return u.Active && u.Email != "" }
