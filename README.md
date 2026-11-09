# Builder Pattern in Go — 內訓範例 repo

10 分鐘內訓簡報「Builder Pattern(Go)」的配套程式碼。簡報上每一段程式碼和「輸出」都在這裡,`go test` 會逐字比對。

**目標行為**:聽完後,看到參數很多的建構函式,能判斷該用 struct literal、Functional Options 還是 Builder,並說出理由。

## 環境與版本

| 項目 | 版本 |
|---|---|
| Go | 1.24(驗證時用 go1.24.7;`errors.Join` 需要 1.20 以上) |
| github.com/Masterminds/squirrel | v1.5.4 |

```bash
git clone https://github.com/holmeslin/builder-pattern-go
cd builder-pattern-go
go vet ./...
go test ./...          # 全部
go test -v ./server/   # 看「忘了 Build() 會編譯失敗」的編譯器輸出
```

## 目錄對照簡報

| 目錄 | 簡報頁 | 情境 | Builder 的理由 | 主要測試 |
|---|---|---|---|---|
| `compare/` | p.3 | 同一個 Server 三種寫法 | — | `Example_threeWays` |
| `docgen/` | p.6 | 文件輸出 | 同一套步驟,多種成品 | `ExampleBuildMonthly`、`TestHTMLBuilderEscapes` |
| `server/` | p.7 | 服務設定 | 統一驗證 | `ExampleBuilder`、`TestBuildReturnsCopy`、`TestForgetBuildDoesNotCompile` |
| `user/` | p.8 | 測試資料 | 預設值覆寫 | `TestAdminCanDelete` 等 |
| `sqlquery/` | p.9 | SQL 查詢 | 有順序的選填子句 | `ExampleListUsers`、`TestClauseOrderIsEnforced`、`TestBuilderIsImmutable` |
| `overkill/` | p.10 | 反例 | 硬套的成本 | `Example_cost`、`TestSameResult` |
| `options/` | p.3、p.11 | Functional Options 對照 | — | `ExampleNew` |
| `cmd/ctorstat/` | p.2 | 掃自己的 repo 拿真實數字 | — | `TestScanSample` |

## 用 ctorstat 拿團隊的真實數字(簡報 p.2)

不用 clone,在自己的 repo 根目錄直接跑:

```bash
go run github.com/holmeslin/builder-pattern-go/cmd/ctorstat@latest -min 5 .
```

列出:參數超過 5 個的 `New*` 函式、`*Builder` 型別、頂層 `With*` 函式。把數字填進簡報 p.2 的 `[__]`。

## 質疑對照表

| # | 質疑 | 回答頁 | 證據 |
|---|---|---|---|
| 1 | 這不就是一堆 setter? | p.7、p.12 | `server.ExampleBuilder`(三個錯誤在 Build() 一次回報)、`TestBuildReturnsCopy` |
| 2 | Go 不是都用 Functional Options? | p.3、p.11、p.12 | `compare.Example_threeWays`、`options.ExampleNew`;Pike 2014、Cheney 2014 |
| 3 | 多一個型別不是過度設計? | p.10、p.12 | `overkill.Example_cost`(多 17 行、5 個函式)、`TestSameResult`(結果一模一樣) |
| 4 | strings.Builder 算 Builder 嗎? | p.12 | p.6 的 `HTMLBuilder` 內部就用 `strings.Builder` 當容器:它是零件,不是 pattern |
| 5 | 共用 builder 會不會互相污染? | p.7、p.9、p.12 | `TestBuildReturnsCopy`、`sqlquery.TestBuilderIsImmutable`;squirrel 是 value receiver(來源 3、4) |
| 6 | 忘了呼叫 Build() 怎麼辦? | p.7、p.12 | `server.TestForgetBuildDoesNotCompile`:`cannot use b (variable of type *server.Builder) as *server.Server value` |
| 7 | Functional Options 不能驗證嗎? | p.3 | `options.ExampleNew`:`Option` 回傳 error,`WithPort(0)` 被拒絕 |
| 8 | HTML 直接拼字串,不怕 XSS? | p.6 | `docgen.TestHTMLBuilderEscapes` |
| 9 | resty、gorm 也跟 squirrel 一樣不可變? | p.4 | 不是。resty `SetHeader` 改自己再回傳自己(來源 6);gorm 只有第一次鏈式呼叫會 clone(來源 7)。形狀是 Builder,不可變只有 squirrel、goqu |

## 來源

1. Gamma, Helm, Johnson, Vlissides. *Design Patterns: Elements of Reusable Object-Oriented Software*. Addison-Wesley, 1994. ISBN 0-201-63361-2. Builder 章節 p.97。Intent:“Separate the construction of a complex object from its representation so that the same construction process can create different representations.”
2. Rob Pike, “Self-referential functions and the design of options”, 2014-01. https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
3. Dave Cheney, “Functional options for friendly APIs”(dotGo 2014 講稿), 2014-10-17. https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
4. squirrel v1.5.4 `SelectBuilder` 是 value receiver、每次回傳新值:https://github.com/Masterminds/squirrel/blob/v1.5.4/select.go#L342-L347 ;ToSql 依固定順序輸出子句:https://github.com/Masterminds/squirrel/blob/v1.5.4/select.go#L96-L141
5. lann/builder(squirrel 的底層):“Package builder provides a method for writing fluent immutable builders.” https://github.com/lann/builder/blob/47ae307949d0/builder.go#L1
6. goqu v9.19.0 每個方法都 `copy()` 出新的 dataset:https://github.com/doug-martin/goqu/blob/v9.19.0/select_dataset.go#L352-L354
7. resty v2.16.5 `SetHeader` 改自己、回傳自己:https://github.com/go-resty/resty/blob/v2.16.5/request.go#L126-L129
8. gorm v1.30.0 分成 chainable / finisher 兩類 API:https://github.com/go-gorm/gorm/blob/v1.30.0/chainable_api.go 、https://github.com/go-gorm/gorm/blob/v1.30.0/finisher_api.go ;`getInstance` 只在 clone 旗標存在時複製:https://github.com/go-gorm/gorm/blob/v1.30.0/gorm.go#L425-L451

> 註:簡報 p.4 的 resty、gorm 兩行另外用本地 replace 編譯驗證過(go vet 通過),沒有放進這個 repo,避免多拉 `golang.org/x/*` 相依。
