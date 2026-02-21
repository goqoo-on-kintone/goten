# goten

kintone REST API 向け Go言語 SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/goqoo-on-kintone/goten.svg)](https://pkg.go.dev/github.com/goqoo-on-kintone/goten)
[![Go Report Card](https://goreportcard.com/badge/github.com/goqoo-on-kintone/goten)](https://goreportcard.com/report/github.com/goqoo-on-kintone/goten)

## 特徴

- **型安全**: Go 1.18+ のジェネリクスを活用した型安全なレコード操作
- **ファサードパターン**: 公式 JS SDK に倣った直感的な API 設計
- **全 API 対応**: Record, App, Space, File, BulkRequest をサポート
- **複数認証方式**: API トークン、パスワード、Basic 認証に対応

## インストール

```bash
go get github.com/goqoo-on-kintone/goten
```

## クイックスタート

```go
package main

import (
    "fmt"
    "os"

    "github.com/goqoo-on-kintone/goten"
    "github.com/goqoo-on-kintone/goten/auth"
    "github.com/goqoo-on-kintone/goten/record"
)

// レコードの型を定義
type MyRecord struct {
    ID struct {
        Value string `json:"value"`
    } `json:"$id"`
    Title struct {
        Value string `json:"value"`
    } `json:"タイトル"`
}

func main() {
    // クライアントを作成
    client := goten.NewClient(goten.ClientConfig{
        BaseURL: "https://your-domain.cybozu.com",
        Auth:    auth.APITokenAuth{Token: os.Getenv("KINTONE_API_TOKEN")},
    })

    // レコードを取得（ジェネリクスで型安全）
    result, err := record.GetRecords[MyRecord](client.Record, record.GetRecordsParams{
        App:   "1",
        Query: "作成日時 > TODAY()",
    })
    if err != nil {
        panic(err)
    }

    for _, rec := range result.Records {
        fmt.Printf("ID: %s, Title: %s\n", rec.ID.Value, rec.Title.Value)
    }
}
```

## API 一覧

### RecordClient

| メソッド | 説明 |
|---------|------|
| `GetRecord[T]` | 単一レコード取得 |
| `GetRecords[T]` | 複数レコード取得 |
| `GetAllRecords[T]` | 全レコード取得（自動ページング） |
| `AddRecord` | レコード追加 |
| `UpdateRecord` | レコード更新 |
| `UpdateRecords` | 複数レコード更新 |
| `DeleteRecords` | レコード削除 |
| `CreateCursor` | カーソル作成 |
| `GetRecordsByCursor[T]` | カーソルでレコード取得 |
| `DeleteCursor` | カーソル削除 |

### AppClient

| メソッド | 説明 |
|---------|------|
| `GetApp` | アプリ情報取得 |
| `GetApps` | 複数アプリ情報取得 |
| `GetFormFields` | フォームフィールド取得 |
| `GetFormLayout` | フォームレイアウト取得 |
| `GetViews` | 一覧設定取得 |

### SpaceClient

| メソッド | 説明 |
|---------|------|
| `GetSpace` | スペース情報取得 |
| `GetSpaceMembers` | スペースメンバー取得 |
| `UpdateSpace` | スペース更新 |
| `UpdateSpaceMembers` | メンバー更新 |
| `AddThread` | スレッド追加 |
| `UpdateThread` | スレッド更新 |
| `AddThreadComment` | コメント追加 |

### FileClient

| メソッド | 説明 |
|---------|------|
| `Upload` | ファイルアップロード |
| `Download` | ファイルダウンロード |

### BulkRequestClient

| メソッド | 説明 |
|---------|------|
| `Send` | バルクリクエスト実行（最大20件） |

## 認証方式

```go
// APIトークン認証
auth.APITokenAuth{Token: "your-api-token"}

// パスワード認証（kintone専用）
auth.PasswordAuth{
    Username: "EXAMPLE",
    Password: "changeme",
}

// Basic認証（プロキシ等で使用）
auth.BasicAuth{
    Username: "EXAMPLE",
    Password: "changeme",
}
```

## ゲストスペース対応

```go
client := goten.NewClient(goten.ClientConfig{
    BaseURL:      "https://your-domain.cybozu.com",
    Auth:         auth.APITokenAuth{Token: "token"},
    GuestSpaceID: intPtr(123),  // ゲストスペースID
})
```

## バルクリクエスト

```go
// Builderを使った便利な構築
builder := bulk.NewBuilder()
builder.
    AddRecord("1", record1).
    AddRecord("1", record2).
    UpdateRecord("1", "100", updates, "")

result, err := client.Bulk.Send(bulk.SendParams{
    Requests: builder.Build(),
})
```

## 開発

```bash
# ビルド
make build

# テスト
make test

# テスト（カバレッジ付き）
make test-coverage

# フォーマット
make fmt

# 静的解析
make lint
```

## ドキュメント

- [設計書](docs/DESIGN.md) - アーキテクチャと設計思想
- [API仕様](docs/API.md) - 公開インターフェース定義
- [ADR](docs/adr/) - Architecture Decision Records

## ライセンス

MIT License

## 関連リンク

- [kintone REST API ドキュメント](https://cybozu.dev/ja/kintone/docs/rest-api/)
- [公式 JavaScript SDK](https://github.com/kintone/js-sdk)
