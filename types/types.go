// Package types は共通の型定義を提供する
package types

// AppID はアプリケーションID
type AppID = string

// RecordID はレコードID
type RecordID = string

// Revision はリビジョン番号
type Revision = string

// FieldValue はフィールド値の基本構造
type FieldValue struct {
	Value any `json:"value"`
}

// UpdateKey は更新キー
type UpdateKey struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
