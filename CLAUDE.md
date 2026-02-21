# CLAUDE.md

## プロジェクト概要

**goten** - kintone REST API向けGo言語SDK。公式JS SDKの設計思想をGoに移植。

## 現在のステータス

v0.1.0 - JS SDK互換APIの実装完了

## パッケージ構成

```
goten/
├── auth/       # 認証（APIトークン、パスワード、Basic）
├── http/       # HTTPクライアント
├── error/      # エラー型定義
├── types/      # 共通型定義
├── record/     # レコードAPI
├── app/        # アプリ設定API
├── space/      # スペースAPI
├── file/       # ファイルAPI
├── bulk/       # バルクリクエストAPI
├── client.go   # ファサード（統合クライアント）
└── examples/   # 使用例
```

## コーディング規約

- コミュニケーション・コードコメント: 日本語
- コンソール出力: 状況に応じて英語/日本語
- フォーマット: gofmt準拠
- Go 1.18+ ジェネリクス使用

## 主要コマンド

```bash
go build ./...    # ビルド
go test ./...     # テスト
go fmt ./...      # フォーマット
```

## 環境変数

```
KINTONE_BASE_URL   # kintoneのベースURL
KINTONE_API_TOKEN  # APIトークン
```

## ドキュメント

- `docs/DESIGN.md` - 設計思想・アーキテクチャ
- `docs/API.md` - 公開APIインターフェース定義
- `docs/adr/` - Architecture Decision Records
- `TODO.md` - 実装状況と将来計画
