// Package sqlquery 是情境 4「SQL 查詢」:用 squirrel 動態組合選填子句。
package sqlquery

import sq "github.com/Masterminds/squirrel"

type Filter struct {
	Status string
	MinAge int
}

// ListUsers 依 Filter 組出查詢;沒給的條件就不加。
func ListUsers(f Filter) (string, []any, error) {
	q := sq.Select("id", "name").From("users")
	if f.Status != "" {
		q = q.Where(sq.Eq{"status": f.Status})
	}
	if f.MinAge > 0 {
		q = q.Where(sq.Gt{"age": f.MinAge})
	}
	return q.OrderBy("name").Limit(20).ToSql() // ToSql 就是 Build()
}
