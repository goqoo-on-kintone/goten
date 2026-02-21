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

// --- アプリ作成/複製API ---

// AddPreviewAppParams はAddPreviewAppのパラメータ
type AddPreviewAppParams struct {
	Name   string // アプリ名（必須）
	Space  string // スペースID（省略可）
	Thread string // スレッドID（省略可）
}

// AddPreviewAppResult はAddPreviewAppの結果
type AddPreviewAppResult struct {
	App      string `json:"app"`
	Revision string `json:"revision"`
}

// CopyAppParams はCopyAppのパラメータ
type CopyAppParams struct {
	App    types.AppID // コピー元アプリID（必須）
	Name   string      // 新しいアプリ名（省略時は元の名前）
	Space  string      // コピー先スペースID（省略可）
	Thread string      // コピー先スレッドID（省略可）
}

// CopyAppResult はCopyAppの結果
type CopyAppResult struct {
	App      string `json:"app"`
	Revision string `json:"revision"`
}

// --- 一覧・レイアウト更新API ---

// UpdateViewsParams はUpdateViewsのパラメータ
type UpdateViewsParams struct {
	App      types.AppID
	Views    map[string]ViewForUpdate
	Revision string
}

// ViewForUpdate は更新用の一覧設定
type ViewForUpdate struct {
	Index      string   `json:"index"`
	Type       string   `json:"type,omitempty"`
	Name       string   `json:"name,omitempty"`
	Fields     []string `json:"fields,omitempty"`
	FilterCond string   `json:"filterCond,omitempty"`
	Sort       string   `json:"sort,omitempty"`
	HTML       string   `json:"html,omitempty"`
	Pager      *bool    `json:"pager,omitempty"`
	Device     string   `json:"device,omitempty"`
}

// UpdateViewsResult はUpdateViewsの結果
type UpdateViewsResult struct {
	Revision string                 `json:"revision"`
	Views    map[string]ViewUpdated `json:"views"`
}

// ViewUpdated は更新後の一覧情報
type ViewUpdated struct {
	ID string `json:"id"`
}

// UpdateFormLayoutParams はUpdateFormLayoutのパラメータ
type UpdateFormLayoutParams struct {
	App      types.AppID
	Layout   []LayoutElement
	Revision string
}

// UpdateFormLayoutResult はUpdateFormLayoutの結果
type UpdateFormLayoutResult struct {
	Revision string `json:"revision"`
}

// --- アプリ設定API ---

// GetAppSettingsParams はGetAppSettingsのパラメータ
type GetAppSettingsParams struct {
	App  types.AppID
	Lang string
}

// AppSettings はアプリの一般設定
type AppSettings struct {
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	Icon                 *AppIcon         `json:"icon,omitempty"`
	Theme                string           `json:"theme,omitempty"`
	TitleField           *TitleField      `json:"titleField,omitempty"`
	EnableThumbnails     *bool            `json:"enableThumbnails,omitempty"`
	EnableBulkDeletion   *bool            `json:"enableBulkDeletion,omitempty"`
	EnableComments       *bool            `json:"enableComments,omitempty"`
	EnableDuplicateRecord *bool           `json:"enableDuplicateRecord,omitempty"`
	NumberPrecision      *NumberPrecision `json:"numberPrecision,omitempty"`
	Revision             string           `json:"revision"`
}

// AppIcon はアプリアイコン
type AppIcon struct {
	Type string `json:"type"` // PRESET or FILE
	Key  string `json:"key"`
}

// TitleField はタイトルフィールド設定
type TitleField struct {
	Code     string `json:"code"`
	Disabled bool   `json:"disabled"`
}

// NumberPrecision は数値精度設定
type NumberPrecision struct {
	Digits    string `json:"digits"`
	RoundType string `json:"roundType"`
}

// GetAppSettingsResult はGetAppSettingsの結果
type GetAppSettingsResult = AppSettings

// UpdateAppSettingsParams はUpdateAppSettingsのパラメータ
type UpdateAppSettingsParams struct {
	App                   types.AppID
	Name                  string
	Description           string
	Icon                  *AppIcon
	Theme                 string
	TitleField            *TitleField
	EnableThumbnails      *bool
	EnableBulkDeletion    *bool
	EnableComments        *bool
	EnableDuplicateRecord *bool
	Revision              string
}

// UpdateAppSettingsResult はUpdateAppSettingsの結果
type UpdateAppSettingsResult struct {
	Revision string `json:"revision"`
}

// --- カスタマイズAPI ---

// GetAppCustomizeParams はGetAppCustomizeのパラメータ
type GetAppCustomizeParams struct {
	App types.AppID
}

// AppCustomize はアプリのカスタマイズ設定
type AppCustomize struct {
	Desktop CustomizeScope `json:"desktop"`
	Mobile  CustomizeScope `json:"mobile"`
	Scope   string         `json:"scope"` // ALL or ADMIN or NONE
}

// CustomizeScope はPC/モバイル別のカスタマイズ設定
type CustomizeScope struct {
	JS  []CustomizeResource `json:"js"`
	CSS []CustomizeResource `json:"css"`
}

// CustomizeResource はJS/CSSリソース
type CustomizeResource struct {
	Type string       `json:"type"` // URL or FILE
	URL  string       `json:"url,omitempty"`
	File *FileInfo    `json:"file,omitempty"`
}

// FileInfo はファイル情報
type FileInfo struct {
	FileKey   string `json:"fileKey"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	ContentType string `json:"contentType,omitempty"`
}

// GetAppCustomizeResult はGetAppCustomizeの結果
type GetAppCustomizeResult struct {
	Desktop  CustomizeScope `json:"desktop"`
	Mobile   CustomizeScope `json:"mobile"`
	Scope    string         `json:"scope"`
	Revision string         `json:"revision"`
}

// UpdateAppCustomizeParams はUpdateAppCustomizeのパラメータ
type UpdateAppCustomizeParams struct {
	App      types.AppID
	Desktop  *CustomizeScopeForUpdate `json:"desktop,omitempty"`
	Mobile   *CustomizeScopeForUpdate `json:"mobile,omitempty"`
	Scope    string                   `json:"scope,omitempty"` // ALL, ADMIN, NONE
	Revision string                   `json:"revision,omitempty"`
}

// CustomizeScopeForUpdate は更新用のカスタマイズ設定
type CustomizeScopeForUpdate struct {
	JS  []CustomizeResourceForUpdate `json:"js,omitempty"`
	CSS []CustomizeResourceForUpdate `json:"css,omitempty"`
}

// CustomizeResourceForUpdate は更新用のJS/CSSリソース
type CustomizeResourceForUpdate struct {
	Type string `json:"type"` // URL or FILE
	URL  string `json:"url,omitempty"`
	File *FileKeyInfo `json:"file,omitempty"`
}

// FileKeyInfo はファイルキー情報
type FileKeyInfo struct {
	FileKey string `json:"fileKey"`
}

// UpdateAppCustomizeResult はUpdateAppCustomizeの結果
type UpdateAppCustomizeResult struct {
	Revision string `json:"revision"`
}

// --- プロセス管理設定API ---

// GetProcessManagementParams はGetProcessManagementのパラメータ
type GetProcessManagementParams struct {
	App  types.AppID
	Lang string // 言語設定（ja, en, zh, default）
}

// ProcessManagement はプロセス管理設定
type ProcessManagement struct {
	Enable   bool              `json:"enable"`
	States   map[string]State  `json:"states"`
	Actions  []Action          `json:"actions"`
	Revision string            `json:"revision"`
}

// State はステータス設定
type State struct {
	Name     string     `json:"name"`
	Index    string     `json:"index"`
	Assignee *Assignee  `json:"assignee,omitempty"`
}

// Assignee は作業者設定
type Assignee struct {
	Type     string         `json:"type"` // ONE, ALL, ANY
	Entities []AssigneeEntity `json:"entities"`
}

// AssigneeEntity は作業者エンティティ
type AssigneeEntity struct {
	Entity EntityInfo `json:"entity"`
	IncludeSubs bool  `json:"includeSubs,omitempty"`
}

// EntityInfo はエンティティ情報
type EntityInfo struct {
	Type string `json:"type"` // USER, GROUP, ORGANIZATION, FIELD_ENTITY, CUSTOM_FIELD, CREATOR
	Code string `json:"code,omitempty"`
}

// Action はアクション設定
type Action struct {
	Name        string `json:"name"`
	From        string `json:"from"`
	To          string `json:"to"`
	FilterCond  string `json:"filterCond,omitempty"`
}

// GetProcessManagementResult はGetProcessManagementの結果
type GetProcessManagementResult = ProcessManagement

// UpdateProcessManagementParams はUpdateProcessManagementのパラメータ
type UpdateProcessManagementParams struct {
	App      types.AppID
	Enable   *bool
	States   map[string]StateForUpdate
	Actions  []ActionForUpdate
	Revision string
}

// StateForUpdate は更新用のステータス設定
type StateForUpdate struct {
	Name     string              `json:"name,omitempty"`
	Index    string              `json:"index,omitempty"`
	Assignee *AssigneeForUpdate  `json:"assignee,omitempty"`
}

// AssigneeForUpdate は更新用の作業者設定
type AssigneeForUpdate struct {
	Type     string               `json:"type"`
	Entities []AssigneeEntityForUpdate `json:"entities,omitempty"`
}

// AssigneeEntityForUpdate は更新用の作業者エンティティ
type AssigneeEntityForUpdate struct {
	Entity      EntityInfo `json:"entity"`
	IncludeSubs bool       `json:"includeSubs,omitempty"`
}

// ActionForUpdate は更新用のアクション設定
type ActionForUpdate struct {
	Name        string `json:"name"`
	From        string `json:"from"`
	To          string `json:"to"`
	FilterCond  string `json:"filterCond,omitempty"`
}

// UpdateProcessManagementResult はUpdateProcessManagementの結果
type UpdateProcessManagementResult struct {
	Revision string `json:"revision"`
}

// --- 権限API ---

// GetAppAclParams はGetAppAclのパラメータ
type GetAppAclParams struct {
	App types.AppID
}

// AppAcl はアプリのアクセス権限
type AppAcl struct {
	Rights   []AppAclRight `json:"rights"`
	Revision string        `json:"revision"`
}

// AppAclRight はアプリのアクセス権限設定
type AppAclRight struct {
	Entity           EntityInfo `json:"entity"`
	IncludeSubs      bool       `json:"includeSubs,omitempty"`
	AppEditable      bool       `json:"appEditable"`
	RecordViewable   bool       `json:"recordViewable"`
	RecordAddable    bool       `json:"recordAddable"`
	RecordEditable   bool       `json:"recordEditable"`
	RecordDeletable  bool       `json:"recordDeletable"`
	RecordImportable bool       `json:"recordImportable"`
	RecordExportable bool       `json:"recordExportable"`
}

// GetAppAclResult はGetAppAclの結果
type GetAppAclResult = AppAcl

// UpdateAppAclParams はUpdateAppAclのパラメータ
type UpdateAppAclParams struct {
	App      types.AppID
	Rights   []AppAclRightForUpdate
	Revision string
}

// AppAclRightForUpdate は更新用のアプリアクセス権限
type AppAclRightForUpdate struct {
	Entity           EntityInfo `json:"entity"`
	IncludeSubs      bool       `json:"includeSubs,omitempty"`
	AppEditable      bool       `json:"appEditable"`
	RecordViewable   bool       `json:"recordViewable"`
	RecordAddable    bool       `json:"recordAddable"`
	RecordEditable   bool       `json:"recordEditable"`
	RecordDeletable  bool       `json:"recordDeletable"`
	RecordImportable bool       `json:"recordImportable"`
	RecordExportable bool       `json:"recordExportable"`
}

// UpdateAppAclResult はUpdateAppAclの結果
type UpdateAppAclResult struct {
	Revision string `json:"revision"`
}

// GetFieldAclParams はGetFieldAclのパラメータ
type GetFieldAclParams struct {
	App types.AppID
}

// FieldAcl はフィールドのアクセス権限
type FieldAcl struct {
	Rights   []FieldAclRight `json:"rights"`
	Revision string          `json:"revision"`
}

// FieldAclRight はフィールドのアクセス権限設定
type FieldAclRight struct {
	Code       string              `json:"code"`
	Entities   []FieldAclEntity    `json:"entities"`
}

// FieldAclEntity はフィールドアクセス権限のエンティティ
type FieldAclEntity struct {
	Accessibility string     `json:"accessibility"` // READ, WRITE, NONE
	Entity        EntityInfo `json:"entity"`
	IncludeSubs   bool       `json:"includeSubs,omitempty"`
}

// GetFieldAclResult はGetFieldAclの結果
type GetFieldAclResult = FieldAcl

// UpdateFieldAclParams はUpdateFieldAclのパラメータ
type UpdateFieldAclParams struct {
	App      types.AppID
	Rights   []FieldAclRightForUpdate
	Revision string
}

// FieldAclRightForUpdate は更新用のフィールドアクセス権限
type FieldAclRightForUpdate struct {
	Code       string                   `json:"code"`
	Entities   []FieldAclEntityForUpdate `json:"entities"`
}

// FieldAclEntityForUpdate は更新用のフィールドアクセス権限エンティティ
type FieldAclEntityForUpdate struct {
	Accessibility string     `json:"accessibility"`
	Entity        EntityInfo `json:"entity"`
	IncludeSubs   bool       `json:"includeSubs,omitempty"`
}

// UpdateFieldAclResult はUpdateFieldAclの結果
type UpdateFieldAclResult struct {
	Revision string `json:"revision"`
}

// GetRecordAclParams はGetRecordAclのパラメータ
type GetRecordAclParams struct {
	App  types.AppID
	Lang string
}

// RecordAcl はレコードのアクセス権限
type RecordAcl struct {
	Rights   []RecordAclRight `json:"rights"`
	Revision string           `json:"revision"`
}

// RecordAclRight はレコードのアクセス権限設定
type RecordAclRight struct {
	FilterCond string              `json:"filterCond"`
	Entities   []RecordAclEntity   `json:"entities"`
}

// RecordAclEntity はレコードアクセス権限のエンティティ
type RecordAclEntity struct {
	Entity      EntityInfo `json:"entity"`
	Viewable    bool       `json:"viewable"`
	Editable    bool       `json:"editable"`
	Deletable   bool       `json:"deletable"`
	IncludeSubs bool       `json:"includeSubs,omitempty"`
}

// GetRecordAclResult はGetRecordAclの結果
type GetRecordAclResult = RecordAcl

// UpdateRecordAclParams はUpdateRecordAclのパラメータ
type UpdateRecordAclParams struct {
	App      types.AppID
	Rights   []RecordAclRightForUpdate
	Revision string
}

// RecordAclRightForUpdate は更新用のレコードアクセス権限
type RecordAclRightForUpdate struct {
	FilterCond string                     `json:"filterCond"`
	Entities   []RecordAclEntityForUpdate `json:"entities"`
}

// RecordAclEntityForUpdate は更新用のレコードアクセス権限エンティティ
type RecordAclEntityForUpdate struct {
	Entity      EntityInfo `json:"entity"`
	Viewable    bool       `json:"viewable"`
	Editable    bool       `json:"editable"`
	Deletable   bool       `json:"deletable"`
	IncludeSubs bool       `json:"includeSubs,omitempty"`
}

// UpdateRecordAclResult はUpdateRecordAclの結果
type UpdateRecordAclResult struct {
	Revision string `json:"revision"`
}
