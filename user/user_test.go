package user

import "testing"

// 測試裡只寫「跟這個測試有關」的差異。
func TestAdminCanDelete(t *testing.T) {
	admin := aUser().Role("admin").Build()
	if !CanDelete(admin) {
		t.Fatal("admin 應該可以刪除")
	}
}

func TestMemberCannotDelete(t *testing.T) {
	if CanDelete(aUser().Build()) {
		t.Fatal("一般會員不應該可以刪除")
	}
}

func TestInactiveCannotLogin(t *testing.T) {
	banned := aUser().Inactive().Build()
	if CanLogin(banned) {
		t.Fatal("停用帳號不應該可以登入")
	}
}

// 預設值本身必須合法,否則每個測試都會被它拖下水。
func TestDefaultUserIsValid(t *testing.T) {
	if !CanLogin(aUser().Build()) {
		t.Fatal("aUser() 的預設值應該是一個能登入的使用者")
	}
}
