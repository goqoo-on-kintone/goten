# API仕様書 - goten

## クライアント初期化

```go
import "github.com/goqoo-on-kintone/goten"

// APIトークン認証
client := goten.NewClient(goten.Options{
    BaseURL:  "https://example.cybozu.com",
    Auth:     goten.APITokenAuth{Token: "your-token"},
})

// パスワード認証
client := goten.NewClient(goten.Options{
    BaseURL:  "https://example.cybozu.com",
    Auth:     goten.PasswordAuth{Username: "user", Password: "pass"},
})

// ゲストスペース
client := goten.NewClient(goten.Options{
    BaseURL:      "https://example.cybozu.com",
    Auth:         goten.APITokenAuth{Token: "your-token"},
    GuestSpaceID: 1,
})
```

---

## RecordClient

### GetRecord

単一レコードを取得する。

```go
func (c *Client) GetRecord[T any](params GetRecordParams) (T, error)

type GetRecordParams struct {
    App AppID
    ID  RecordID
}
```

**使用例:**
```go
record, err := client.Record.GetRecord[MyRecord](record.GetRecordParams{
    App: "1",
    ID:  "100",
})
```

### GetRecords

複数レコードを取得する（最大500件）。

```go
func (c *Client) GetRecords[T any](params GetRecordsParams) (GetRecordsResult[T], error)

type GetRecordsParams struct {
    App        AppID
    Fields     []string  // 取得するフィールド（省略時は全フィールド）
    Query      string    // クエリ文字列
    TotalCount bool      // 総件数を取得するか
}

type GetRecordsResult[T any] struct {
    Records    []T
    TotalCount *string
}
```

### GetAllRecords

全レコードを取得する（ページング自動処理）。

```go
func (c *Client) GetAllRecords[T any](params GetAllRecordsParams) ([]T, error)

type GetAllRecordsParams struct {
    App       AppID
    Fields    []string
    Condition string   // WHERE句相当
    OrderBy   string   // ORDER BY句相当
}
```

### AddRecord

レコードを1件追加する。

```go
func (c *Client) AddRecord(params AddRecordParams) (AddRecordResult, error)

type AddRecordParams struct {
    App    AppID
    Record map[string]FieldValue
}

type AddRecordResult struct {
    ID       string
    Revision string
}
```

### AddRecords

レコードを複数追加する（最大100件）。

```go
func (c *Client) AddRecords(params AddRecordsParams) (AddRecordsResult, error)

type AddRecordsParams struct {
    App     AppID
    Records []map[string]FieldValue
}

type AddRecordsResult struct {
    IDs       []string
    Revisions []string
}
```

### UpdateRecord

レコードを1件更新する。

```go
func (c *Client) UpdateRecord(params UpdateRecordParams) (UpdateRecordResult, error)

type UpdateRecordParams struct {
    App      AppID
    ID       RecordID           // IDまたはUpdateKeyのどちらか必須
    UpdateKey *UpdateKey
    Record   map[string]FieldValue
    Revision *Revision          // 省略時はリビジョンチェックなし
}

type UpdateKey struct {
    Field string
    Value string
}
```

### UpdateRecords

レコードを複数更新する（最大100件）。

```go
func (c *Client) UpdateRecords(params UpdateRecordsParams) (UpdateRecordsResult, error)
```

### DeleteRecords

レコードを削除する（最大100件）。

```go
func (c *Client) DeleteRecords(params DeleteRecordsParams) error

type DeleteRecordsParams struct {
    App       AppID
    IDs       []RecordID
    Revisions []Revision  // 省略可
}
```

### カーソルAPI

```go
func (c *Client) CreateCursor(params CreateCursorParams) (CreateCursorResult, error)
func (c *Client) GetRecordsByCursor[T any](params GetRecordsByCursorParams) (GetRecordsByCursorResult[T], error)
func (c *Client) DeleteCursor(params DeleteCursorParams) error
```

---

## FileClient

### Upload

ファイルをアップロードする。

```go
func (c *Client) Upload(params UploadParams) (UploadResult, error)

type UploadParams struct {
    File     io.Reader
    FileName string
}

type UploadResult struct {
    FileKey string
}
```

### Download

ファイルをダウンロードする。

```go
func (c *Client) Download(params DownloadParams) (io.ReadCloser, error)

type DownloadParams struct {
    FileKey string
}
```

---

## AppClient

### GetFormFields

フィールド情報を取得する。

```go
func (c *Client) GetFormFields(params GetFormFieldsParams) (GetFormFieldsResult, error)

type GetFormFieldsParams struct {
    App     AppID
    Preview bool  // プレビュー環境から取得
}
```

### GetApp

アプリ情報を取得する。

```go
func (c *Client) GetApp(params GetAppParams) (App, error)
```

---

## 型定義

### 基本型

```go
type AppID = string | int
type RecordID = string | int
type Revision = string | int
```

### フィールド値

```go
type FieldValue struct {
    Value any `json:"value"`
}

// 型付きフィールド（gotenks/types）
type SingleLineTextField struct {
    Type  string `json:"type"`
    Value string `json:"value"`
}
```

---

## エラー型

```go
type KintoneRestAPIError struct {
    Status  int
    Code    string
    Message string
    ID      string
    Errors  map[string]any
}

type KintoneAllRecordsError struct {
    ProcessedRecords   any
    UnprocessedRecords []any
    Error              KintoneRestAPIError
}
```
