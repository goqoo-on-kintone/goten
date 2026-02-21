# ADR-002: レコード取得におけるジェネリクスの採用

## ステータス

承認

## コンテキスト

kintoneのレコードは、アプリごとに異なるフィールド構造を持つ。公式Go SDKでは `map[string]interface{}` を使用しているため、型安全性がなく、フィールドアクセス時に型アサーションが必要となる。

```go
// 公式Go SDK
record, _ := app.GetRecord(1)
name := record.Fields["名前"].(SingleLineTextField).String()  // 実行時エラーの可能性
```

## 決定

Go 1.18以降のジェネリクスを活用し、ユーザー定義の構造体でレコードを取得できるようにする。

```go
func GetRecords[T any](c *Client, params GetRecordsParams) ([]T, error)
```

ユーザーは以下のように型安全にレコードを扱える：

```go
type MyRecord struct {
    Name types.SingleLineTextField `json:"名前"`
    Age  types.NumberField         `json:"年齢"`
}

records, err := client.Record.GetRecords[MyRecord](params)
fmt.Println(records[0].Name.Value)  // コンパイル時に型チェック
```

## 選択肢

1. **map[string]interface{}**（公式Go SDK方式）
   - メリット: Go 1.18以前でも動作
   - デメリット: 型安全性なし

2. **ジェネリクス**
   - メリット: 型安全、IDE補完
   - デメリット: Go 1.18以降が必須

3. **コード生成**
   - メリット: 型安全
   - デメリット: ビルドプロセスが複雑

## 結果

- コンパイル時に型エラーを検出可能
- IDEによるコード補完が効く
- `gotenks/types` パッケージと連携し、フィールド型を再利用
- Go 1.18未満の環境ではビルド不可（許容範囲）
