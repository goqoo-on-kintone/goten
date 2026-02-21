// Package app はアプリ設定APIを提供する
package app

import "github.com/goqoo-on-kintone/goten/types"

// GetAppParams はGetAppのパラメータ
type GetAppParams struct {
	App types.AppID
}

// App はアプリの基本情報
type App struct {
	AppID       string `json:"appId"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SpaceID     string `json:"spaceId,omitempty"`
	ThreadID    string `json:"threadId,omitempty"`
	CreatedAt   string `json:"createdAt"`
	Creator     User   `json:"creator"`
	ModifiedAt  string `json:"modifiedAt"`
	Modifier    User   `json:"modifier"`
}

// User はユーザー情報
type User struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetAppsParams はGetAppsのパラメータ
type GetAppsParams struct {
	IDs      []types.AppID // アプリID一覧
	Codes    []string      // アプリコード一覧
	Name     string        // アプリ名（部分一致）
	SpaceIDs []string      // スペースID一覧
	Limit    int           // 取得件数（最大100）
	Offset   int           // オフセット
}

// GetAppsResult はGetAppsの結果
type GetAppsResult struct {
	Apps []App `json:"apps"`
}

// GetFormFieldsParams はGetFormFieldsのパラメータ
type GetFormFieldsParams struct {
	App  types.AppID
	Lang string // 言語設定（ja, en, zh, default）
}

// FieldProperty はフィールドプロパティ
type FieldProperty struct {
	Type         string            `json:"type"`
	Code         string            `json:"code"`
	Label        string            `json:"label"`
	NoLabel      bool              `json:"noLabel,omitempty"`
	Required     bool              `json:"required,omitempty"`
	Unique       bool              `json:"unique,omitempty"`
	MaxValue     string            `json:"maxValue,omitempty"`
	MinValue     string            `json:"minValue,omitempty"`
	MaxLength    string            `json:"maxLength,omitempty"`
	MinLength    string            `json:"minLength,omitempty"`
	DefaultValue any               `json:"defaultValue,omitempty"`
	Options      map[string]Option `json:"options,omitempty"`
	// 他にも多数のプロパティがあるが、必要に応じて追加
}

// Option はドロップダウンなどの選択肢
type Option struct {
	Label string `json:"label"`
	Index string `json:"index"`
}

// GetFormFieldsResult はGetFormFieldsの結果
type GetFormFieldsResult struct {
	Properties map[string]FieldProperty `json:"properties"`
	Revision   string                   `json:"revision"`
}

// GetFormLayoutParams はGetFormLayoutのパラメータ
type GetFormLayoutParams struct {
	App types.AppID
}

// LayoutElement はレイアウト要素
type LayoutElement struct {
	Type   string          `json:"type"`
	Code   string          `json:"code,omitempty"`
	Fields []LayoutField   `json:"fields,omitempty"`
	Layout []LayoutElement `json:"layout,omitempty"` // GROUP, SUBTABLEの場合
}

// LayoutField はレイアウト内のフィールド
type LayoutField struct {
	Type      string `json:"type"`
	Code      string `json:"code,omitempty"`
	Label     string `json:"label,omitempty"`
	ElementID string `json:"elementId,omitempty"`
	Size      *Size  `json:"size,omitempty"`
}

// Size はフィールドサイズ
type Size struct {
	Width       string `json:"width,omitempty"`
	Height      string `json:"height,omitempty"`
	InnerHeight string `json:"innerHeight,omitempty"`
}

// GetFormLayoutResult はGetFormLayoutの結果
type GetFormLayoutResult struct {
	Layout   []LayoutElement `json:"layout"`
	Revision string          `json:"revision"`
}

// GetViewsParams はGetViewsのパラメータ
type GetViewsParams struct {
	App  types.AppID
	Lang string // 言語設定
}

// View は一覧の設定
type View struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"` // LIST, CALENDAR, CUSTOM
	Name        string   `json:"name"`
	BuiltinType string   `json:"builtinType,omitempty"`
	Fields      []string `json:"fields,omitempty"`
	FilterCond  string   `json:"filterCond,omitempty"`
	Sort        string   `json:"sort,omitempty"`
	Index       string   `json:"index"`
	HTML        string   `json:"html,omitempty"` // CUSTOMの場合
	Pager       bool     `json:"pager,omitempty"`
	Device      string   `json:"device,omitempty"` // ANY, DESKTOP, MOBILE
}

// GetViewsResult はGetViewsの結果
type GetViewsResult struct {
	Views    map[string]View `json:"views"`
	Revision string          `json:"revision"`
}

// UpdateFormFieldsParams はUpdateFormFieldsのパラメータ
type UpdateFormFieldsParams struct {
	App        types.AppID
	Properties map[string]FieldProperty // 更新するフィールド設定
	Revision   string                   // リビジョン（省略可）
}

// UpdateFormFieldsResult はUpdateFormFieldsの結果
type UpdateFormFieldsResult struct {
	Revision string `json:"revision"`
}

// AddFormFieldsParams はAddFormFieldsのパラメータ
type AddFormFieldsParams struct {
	App        types.AppID
	Properties map[string]FieldProperty // 追加するフィールド設定
	Revision   string                   // リビジョン（省略可）
}

// AddFormFieldsResult はAddFormFieldsの結果
type AddFormFieldsResult struct {
	Revision string `json:"revision"`
}

// DeleteFormFieldsParams はDeleteFormFieldsのパラメータ
type DeleteFormFieldsParams struct {
	App      types.AppID
	Fields   []string // 削除するフィールドコード
	Revision string   // リビジョン（省略可）
}

// DeleteFormFieldsResult はDeleteFormFieldsの結果
type DeleteFormFieldsResult struct {
	Revision string `json:"revision"`
}

// DeployAppParams はDeployAppのパラメータ
type DeployAppParams struct {
	Apps   []DeployAppItem // デプロイするアプリ
	Revert bool            // true: 変更を破棄
}

// DeployAppItem はデプロイ対象のアプリ
type DeployAppItem struct {
	App      types.AppID `json:"app"`
	Revision string      `json:"revision,omitempty"`
}

// GetDeployStatusParams はGetDeployStatusのパラメータ
type GetDeployStatusParams struct {
	Apps []types.AppID
}

// GetDeployStatusResult はGetDeployStatusの結果
type GetDeployStatusResult struct {
	Apps []DeployStatus `json:"apps"`
}

// DeployStatus はデプロイステータス
type DeployStatus struct {
	App    string `json:"app"`
	Status string `json:"status"` // PROCESSING, SUCCESS, FAIL, CANCEL
}
