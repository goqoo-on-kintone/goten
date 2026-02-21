# ADR-001: ファサードパターンの採用

## ステータス

承認

## コンテキスト

kintone REST APIは多岐にわたる機能（レコード操作、アプリ設定、スペース管理など）を持つ。公式Go SDKは単一の`App`構造体にすべてのメソッドを集約しているが、これは以下の問題を引き起こす：

- コードの見通しが悪い
- 関心の分離ができていない
- テストが困難

## 決定

公式JS SDKに倣い、ファサードパターンを採用する。

```go
type KintoneRestAPIClient struct {
    Record *record.Client
    App    *app.Client
    Space  *space.Client
    File   *file.Client
}
```

ユーザーは `client.Record.GetRecords()` のように、機能ごとにグループ化されたAPIにアクセスする。

## 選択肢

1. **単一構造体方式**（公式Go SDK方式）
   - メリット: シンプル
   - デメリット: 肥大化、関心の分離なし

2. **ファサードパターン**（公式JS SDK方式）
   - メリット: 関心の分離、テスタビリティ向上
   - デメリット: 構造がやや複雑

3. **関数ベースAPI**
   - メリット: シンプル
   - デメリット: クライアント状態の管理が煩雑

## 結果

- 各クライアント（RecordClient, AppClientなど）を独立して開発・テスト可能
- ユーザーはJS SDKと似たAPIで直感的に使用可能
- パッケージ構造が明確になる
