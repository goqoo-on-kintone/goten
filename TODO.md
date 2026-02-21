# TODO

## 実装済み機能 (v0.1.0)

### 基盤
- [x] 認証モジュール（APIトークン、パスワード、Basic認証）
- [x] HTTPクライアント抽象化
- [x] エラー型定義
- [x] context.Context対応

### Record API
- [x] GetRecord / GetRecords / GetAllRecords
- [x] AddRecord / AddRecords
- [x] UpdateRecord / UpdateRecords
- [x] DeleteRecords
- [x] CreateCursor / GetRecordsByCursor / DeleteCursor
- [x] GetRecordComments / AddRecordComment / DeleteRecordComment
- [x] UpdateRecordStatus / UpdateRecordsStatus
- [x] UpsertRecord

### App API
- [x] GetApp / GetApps
- [x] GetFormFields / AddFormFields / UpdateFormFields / DeleteFormFields
- [x] GetFormLayout / UpdateFormLayout
- [x] GetViews / UpdateViews
- [x] GetAppSettings / UpdateAppSettings
- [x] GetAppCustomize / UpdateAppCustomize
- [x] GetProcessManagement / UpdateProcessManagement
- [x] GetAppAcl / UpdateAppAcl
- [x] GetFieldAcl / UpdateFieldAcl
- [x] GetRecordAcl / UpdateRecordAcl
- [x] AddPreviewApp / CopyApp
- [x] DeployApp / GetDeployStatus

### Space API
- [x] GetSpace / UpdateSpace
- [x] GetSpaceMembers / UpdateSpaceMembers
- [x] AddThread / UpdateThread / AddThreadComment
- [x] DeleteSpace
- [x] AddGuests / AddGuestsToSpace / UpdateSpaceGuests / DeleteGuests

### File API
- [x] Upload / Download

### Bulk API
- [x] Send (BulkRequest)
- [x] Builder パターン

---

## 将来実装予定（kintone REST APIにあってJS SDKにないもの）

### レコード関連
| API | エンドポイント | 説明 | 優先度 |
|-----|---------------|------|--------|
| 作業者更新 | `PUT /k/v1/record/assignees.json` | ステータス変更なしで作業者のみ更新 | 中 |
| アクセス権評価 | `GET /k/v1/records/acl/evaluate.json` | レコードのアクセス権を評価 | 低 |

### アプリ設定関連
| API | エンドポイント | 説明 | 優先度 |
|-----|---------------|------|--------|
| グラフ設定 | `GET/PUT app/reports.json` | グラフ・集計設定 | 高 |
| 通知設定 | `notifications/general, perRecord, reminder` | 各種通知設定 | 高 |
| アクション設定 | `GET/PUT actions.json` | レコードアクション | 中 |
| プラグイン設定 | `GET/POST app/plugins.json` | アプリのプラグイン | 中 |
| 管理者用メモ | `GET/PUT adminNotes.json` | 管理者向けメモ | 低 |
| 使用状況取得 | `GET apps/statistics.json` | アプリの使用統計 | 低 |
| アプリ移動 | `POST app/move.json` | スペース間移動 | 低 |

### スペース関連
| API | エンドポイント | 説明 | 優先度 |
|-----|---------------|------|--------|
| テンプレートから作成 | `POST template/space.json` | テンプレートでスペース作成 | 中 |
| 本文更新 | `PUT space/body.json` | スペース本文のみ更新 | 低 |
| 使用状況取得 | `GET space/statistics.json` | スペースの使用統計 | 低 |

### プラグイン管理（システム管理者向け）
| API | エンドポイント | 説明 | 優先度 |
|-----|---------------|------|--------|
| 一覧取得 | `GET plugins.json` | インストール済みプラグイン | 低 |
| 必須プラグイン | `GET plugins/required.json` | 必須プラグイン一覧 | 低 |
| 使用アプリ取得 | `GET plugin/apps.json` | プラグイン使用アプリ | 低 |
| インストール等 | `POST/PUT/DELETE plugin.json` | プラグイン管理 | 低 |

### その他
| API | エンドポイント | 説明 | 優先度 |
|-----|---------------|------|--------|
| API一覧取得 | `GET apis.json` | REST API一覧 | 低 |

---

## その他の改善項目

| 項目 | 説明 | 優先度 |
|------|------|--------|
| リトライ/レート制限 | 自動リトライ、指数バックオフ | 中 |
| GoDocコメント充実 | pkg.go.dev向け | 中 |
| インテグレーションテスト | 実際のkintone環境でのテスト | 中 |
| CI/CD設定 | GitHub Actions | 高 |

---

## マイルストーン

| バージョン | 目標 |
|-----------|------|
| v0.1.0 | JS SDK互換API完了（現在） |
| v0.2.0 | グラフ設定・通知設定API追加 |
| v0.3.0 | 残りのkintone REST API対応 |
| v1.0.0 | 安定版リリース |
