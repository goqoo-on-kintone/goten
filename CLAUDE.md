# CLAUDE.md

## プロジェクト概要

**goten** - kintone REST API向けGo言語SDK。公式JS SDKの設計思想をGoに移植。

## コーディング規約

- コミュニケーション・コードコメント: 日本語
- コンソール出力: 状況に応じて英語/日本語
- フォーマット: gofmt準拠

## ドキュメント構成

- `docs/DESIGN.md` - 設計思想・アーキテクチャ
- `docs/API.md` - 公開APIインターフェース定義
- `docs/adr/` - Architecture Decision Records

## 主要コマンド

```bash
go build -o kintone-client  # ビルド
go test ./...               # テスト
```

## 環境変数

```
KINTONE_API_TOKEN  # APIトークン
```
