// Package error はkintone APIエラー型を提供する
package error

import "fmt"

// KintoneRestAPIError はkintone REST APIエラー
type KintoneRestAPIError struct {
	Status  int            // HTTPステータスコード
	Code    string         // kintoneエラーコード
	Message string         // エラーメッセージ
	ID      string         // エラーID
	Errors  map[string]any // 詳細エラー情報
}

// Error はerrorインターフェースを実装
func (e *KintoneRestAPIError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s (%s)", e.Status, e.Code, e.Message, e.ID)
}

// KintoneAllRecordsError は大量レコード処理時のエラー
type KintoneAllRecordsError struct {
	ProcessedRecords   any                  // 処理済みレコード
	UnprocessedRecords []any                // 未処理レコード
	Error              *KintoneRestAPIError // 原因のエラー
	ErrorIndex         int                  // エラーが発生したインデックス
}
