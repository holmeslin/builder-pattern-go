package sqlquery

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
)

func ExampleListUsers() {
	fmt.Println(ListUsers(Filter{Status: "active", MinAge: 18}))
	fmt.Println(ListUsers(Filter{}))
	// Output:
	// SELECT id, name FROM users WHERE status = ? AND age > ? ORDER BY name LIMIT 20 [active 18] <nil>
	// SELECT id, name FROM users ORDER BY name LIMIT 20 [] <nil>
}

// 呼叫順序亂寫,輸出的子句順序仍然合法:順序由 builder 保證。
func TestClauseOrderIsEnforced(t *testing.T) {
	got, _, _ := sq.Select("id").Limit(20).
		Where(sq.Eq{"a": 1}).
		From("users").ToSql()
	want := "SELECT id FROM users WHERE a = ? LIMIT 20"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

// squirrel 的方法是 value receiver,每次回傳新的 builder:
// 同一個 base 分支出兩個查詢,互不影響。
func TestBuilderIsImmutable(t *testing.T) {
	base := sq.Select("*").From("users")
	a, _, _ := base.Where(sq.Eq{"x": 1}).ToSql()
	b, _, _ := base.OrderBy("name").ToSql()
	if a != "SELECT * FROM users WHERE x = ?" {
		t.Fatalf("a = %q", a)
	}
	if b != "SELECT * FROM users ORDER BY name" {
		t.Fatalf("b 被 a 的 Where 污染了: %q", b)
	}
}
