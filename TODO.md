# TODO

タスク管理はGitHub Issuesに移行予定です。

## 初期Issue案（ラベル付き）

### Phase 1: 基盤整備

| Issue | ラベル |
|-------|--------|
| 認証モジュール実装（auth/） | `enhancement`, `core` |
| HTTPクライアント抽象化（http/） | `enhancement`, `core` |
| エラー型定義（error/） | `enhancement`, `core` |
| プロジェクト構造リファクタリング | `refactor`, `core` |

### Phase 2: コアクライアント

| Issue | ラベル |
|-------|--------|
| RecordClient実装 | `enhancement`, `api` |
| FileClient実装 | `enhancement`, `api` |
| ファサード（KintoneRestAPIClient）実装 | `enhancement`, `core` |

### Phase 3: 拡張機能

| Issue | ラベル |
|-------|--------|
| AppClient実装 | `enhancement`, `api` |
| SpaceClient実装 | `enhancement`, `api` |
| BulkRequestClient実装 | `enhancement`, `api` |

### Phase 4: ドキュメント・テスト

| Issue | ラベル |
|-------|--------|
| README.md作成 | `documentation` |
| ユニットテスト追加 | `test` |
| 使用例（examples/）追加 | `documentation` |

---

## GitHub Issues作成コマンド

```bash
# リポジトリ初期化後に実行
gh label create core --color 0366d6 --description "コア機能"
gh label create api --color 1d76db --description "API実装"
gh label create refactor --color fbca04 --description "リファクタリング"

gh issue create --title "認証モジュール実装" --label "enhancement,core"
gh issue create --title "HTTPクライアント抽象化" --label "enhancement,core"
# ...
```

---

## マイルストーン案

| マイルストーン | 目標 |
|---------------|------|
| v0.1.0 | RecordClient基本機能（GetRecords, AddRecord） |
| v0.2.0 | RecordClient完成 + FileClient |
| v0.3.0 | AppClient + SpaceClient |
| v1.0.0 | 全API対応 + ドキュメント完備 |
