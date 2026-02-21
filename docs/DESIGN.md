# 設計書 - goten

## 背景

kintone公式Go SDK (`go-kintone`) は機能が限定的で、ジェネリクスによる型安全なレコード操作ができない。一方、公式JS SDK (`@kintone/rest-api-client`) は型システムを活用した優れた設計を持つ。本プロジェクトは、JS SDKの設計思想をGo言語に移植することを目指す。

## 参考SDK比較

### 公式Go SDK (`kintone-labs/go-kintone`)

| 項目 | 内容 |
|------|------|
| 構造 | 単一パッケージ、`app.go`に大半のロジック |
| 型定義 | `map[string]interface{}`ベース |
| ジェネリクス | 未使用（Go 1.18以前互換） |
| 認証 | APIトークン、パスワード、Basic認証 |
| 対応API | レコードCRUD、コメント、ファイル、カーソル |

### 公式JS SDK (`@kintone/rest-api-client`)

| 項目 | 内容 |
|------|------|
| 構造 | ファサードパターン、機能別クライアント分割 |
| 型定義 | ジェネリクス、Discriminated Union、Mapped Types |
| 認証 | APIトークン、パスワード、OAuth、セッション、証明書 |
| 対応API | ほぼ全API（レコード、アプリ設定、スペース、プラグイン、バルク） |
| 特徴 | HTTP層抽象化、プラットフォーム依存分離 |

## アーキテクチャ

### ディレクトリ構造

```
goten/
├── client.go                 # KintoneRestAPIClient（ファサード）
├── record/
│   ├── client.go            # RecordClient
│   └── types.go             # レコード関連型
├── app/
│   ├── client.go            # AppClient
│   └── types.go
├── space/
│   └── client.go            # SpaceClient
├── file/
│   └── client.go            # FileClient
├── bulk/
│   └── client.go            # BulkRequestClient
├── http/
│   ├── client.go            # HTTPクライアント抽象化
│   └── config.go            # リクエスト設定
├── auth/
│   └── auth.go              # 認証インターフェース
├── error/
│   └── error.go             # カスタムエラー
└── types/
    └── field.go             # フィールド型定義
```

### 設計パターン

#### 1. ファサードパターン

```go
type KintoneRestAPIClient struct {
    Record *record.Client
    App    *app.Client
    Space  *space.Client
    File   *file.Client
}

// 使用例
client := goten.NewClient(options)
records, err := client.Record.GetRecords[MyRecord](params)
```

#### 2. ジェネリクス対応

```go
func (c *Client) GetRecords[T any](params GetRecordsParams) ([]T, error)
func (c *Client) GetRecord[T any](params GetRecordParams) (T, error)
```

#### 3. 認証の抽象化

```go
type Auth interface {
    Apply(req *http.Request)
}

type APITokenAuth struct { Token string }
type PasswordAuth struct { Username, Password string }
```

#### 4. HTTP層の抽象化

```go
type HttpClient interface {
    Get(path string, params any) (*http.Response, error)
    Post(path string, body any) (*http.Response, error)
    Put(path string, body any) (*http.Response, error)
    Delete(path string, params any) (*http.Response, error)
}
```

## API制限値

| API | 上限 |
|-----|------|
| getRecords | 500件 |
| addRecords | 100件 |
| updateRecords | 100件 |
| deleteRecords | 100件 |
| bulkRequest | 20リクエスト |

## 依存パッケージ

- `github.com/joho/godotenv` - 環境変数読み込み
- `github.com/goqoo-on-kintone/gotenks/types` - フィールド型定義（外部）
