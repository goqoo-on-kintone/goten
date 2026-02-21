// Package app はアプリ設定APIを提供する
package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goqoo-on-kintone/goten/http"
)

// Client はアプリ設定クライアント
type Client struct {
	httpClient *http.DefaultClient
}

// NewClient は新しいAppClientを作成する
func NewClient(httpClient *http.DefaultClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// GetApp はアプリの設定を取得する
func (c *Client) GetApp(ctx context.Context, params GetAppParams) (*App, error) {
	reqBody := map[string]any{
		"id": params.App,
	}

	body, err := c.httpClient.GetWithBody(ctx, "app", reqBody)
	if err != nil {
		return nil, err
	}

	var result App
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetApps は複数アプリの設定を取得する
func (c *Client) GetApps(ctx context.Context, params GetAppsParams) (*GetAppsResult, error) {
	reqBody := map[string]any{}

	if len(params.IDs) > 0 {
		reqBody["ids"] = params.IDs
	}
	if len(params.Codes) > 0 {
		reqBody["codes"] = params.Codes
	}
	if params.Name != "" {
		reqBody["name"] = params.Name
	}
	if len(params.SpaceIDs) > 0 {
		reqBody["spaceIds"] = params.SpaceIDs
	}
	if params.Limit > 0 {
		reqBody["limit"] = params.Limit
	}
	if params.Offset > 0 {
		reqBody["offset"] = params.Offset
	}

	body, err := c.httpClient.GetWithBody(ctx, "apps", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetAppsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetFormFields はフォームのフィールド設定を取得する
func (c *Client) GetFormFields(ctx context.Context, params GetFormFieldsParams) (*GetFormFieldsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/form/fields", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetFormFieldsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetFormLayout はフォームのレイアウト設定を取得する
func (c *Client) GetFormLayout(ctx context.Context, params GetFormLayoutParams) (*GetFormLayoutResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/form/layout", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetFormLayoutResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetViews は一覧の設定を取得する
func (c *Client) GetViews(ctx context.Context, params GetViewsParams) (*GetViewsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/views", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetViewsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateFormFields はフォームのフィールド設定を更新する（プレビュー環境）
func (c *Client) UpdateFormFields(ctx context.Context, params UpdateFormFieldsParams) (*UpdateFormFieldsResult, error) {
	reqBody := map[string]any{
		"app":        params.App,
		"properties": params.Properties,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/form/fields", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateFormFieldsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// AddFormFields はフォームにフィールドを追加する（プレビュー環境）
func (c *Client) AddFormFields(ctx context.Context, params AddFormFieldsParams) (*AddFormFieldsResult, error) {
	reqBody := map[string]any{
		"app":        params.App,
		"properties": params.Properties,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Post(ctx, "preview/app/form/fields", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddFormFieldsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// DeleteFormFields はフォームのフィールドを削除する（プレビュー環境）
func (c *Client) DeleteFormFields(ctx context.Context, params DeleteFormFieldsParams) (*DeleteFormFieldsResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"fields": params.Fields,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.DeleteWithBody(ctx, "preview/app/form/fields", reqBody)
	if err != nil {
		return nil, err
	}

	var result DeleteFormFieldsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// DeployApp はアプリの設定を運用環境に反映する
func (c *Client) DeployApp(ctx context.Context, params DeployAppParams) error {
	reqBody := map[string]any{
		"apps": params.Apps,
	}

	if params.Revert {
		reqBody["revert"] = true
	}

	_, err := c.httpClient.Post(ctx, "preview/app/deploy", reqBody)
	return err
}

// GetDeployStatus はデプロイのステータスを取得する
func (c *Client) GetDeployStatus(ctx context.Context, params GetDeployStatusParams) (*GetDeployStatusResult, error) {
	reqBody := map[string]any{
		"apps": params.Apps,
	}

	body, err := c.httpClient.GetWithBody(ctx, "preview/app/deploy", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetDeployStatusResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- アプリ作成/複製API ---

// AddPreviewApp は新しいアプリを作成する（プレビュー環境）
// 作成後にDeployAppで運用環境に反映する必要がある
func (c *Client) AddPreviewApp(ctx context.Context, params AddPreviewAppParams) (*AddPreviewAppResult, error) {
	reqBody := map[string]any{
		"name": params.Name,
	}

	if params.Space != "" {
		reqBody["space"] = params.Space
	}
	if params.Thread != "" {
		reqBody["thread"] = params.Thread
	}

	body, err := c.httpClient.Post(ctx, "preview/app", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddPreviewAppResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// CopyApp はアプリを複製する
// 複製されたアプリは自動的に運用環境に反映される
func (c *Client) CopyApp(ctx context.Context, params CopyAppParams) (*CopyAppResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Name != "" {
		reqBody["name"] = params.Name
	}
	if params.Space != "" {
		reqBody["space"] = params.Space
	}
	if params.Thread != "" {
		reqBody["thread"] = params.Thread
	}

	body, err := c.httpClient.Post(ctx, "app", reqBody)
	if err != nil {
		return nil, err
	}

	var result CopyAppResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- 一覧・レイアウト更新API ---

// UpdateViews は一覧の設定を更新する（プレビュー環境）
func (c *Client) UpdateViews(ctx context.Context, params UpdateViewsParams) (*UpdateViewsResult, error) {
	reqBody := map[string]any{
		"app":   params.App,
		"views": params.Views,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/views", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateViewsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateFormLayout はフォームレイアウトを更新する（プレビュー環境）
func (c *Client) UpdateFormLayout(ctx context.Context, params UpdateFormLayoutParams) (*UpdateFormLayoutResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"layout": params.Layout,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/form/layout", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateFormLayoutResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- アプリ設定API ---

// GetAppSettings はアプリの一般設定を取得する
func (c *Client) GetAppSettings(ctx context.Context, params GetAppSettingsParams) (*GetAppSettingsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/settings", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetAppSettingsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateAppSettings はアプリの一般設定を更新する（プレビュー環境）
func (c *Client) UpdateAppSettings(ctx context.Context, params UpdateAppSettingsParams) (*UpdateAppSettingsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Name != "" {
		reqBody["name"] = params.Name
	}
	if params.Description != "" {
		reqBody["description"] = params.Description
	}
	if params.Icon != nil {
		reqBody["icon"] = params.Icon
	}
	if params.Theme != "" {
		reqBody["theme"] = params.Theme
	}
	if params.TitleField != nil {
		reqBody["titleField"] = params.TitleField
	}
	if params.EnableThumbnails != nil {
		reqBody["enableThumbnails"] = *params.EnableThumbnails
	}
	if params.EnableBulkDeletion != nil {
		reqBody["enableBulkDeletion"] = *params.EnableBulkDeletion
	}
	if params.EnableComments != nil {
		reqBody["enableComments"] = *params.EnableComments
	}
	if params.EnableDuplicateRecord != nil {
		reqBody["enableDuplicateRecord"] = *params.EnableDuplicateRecord
	}
	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/settings", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateAppSettingsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- カスタマイズAPI ---

// GetAppCustomize はアプリのJavaScript/CSSカスタマイズ設定を取得する
func (c *Client) GetAppCustomize(ctx context.Context, params GetAppCustomizeParams) (*GetAppCustomizeResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/customize", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetAppCustomizeResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateAppCustomize はアプリのJavaScript/CSSカスタマイズ設定を更新する（プレビュー環境）
func (c *Client) UpdateAppCustomize(ctx context.Context, params UpdateAppCustomizeParams) (*UpdateAppCustomizeResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Desktop != nil {
		reqBody["desktop"] = params.Desktop
	}
	if params.Mobile != nil {
		reqBody["mobile"] = params.Mobile
	}
	if params.Scope != "" {
		reqBody["scope"] = params.Scope
	}
	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/customize", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateAppCustomizeResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- プロセス管理設定API ---

// GetProcessManagement はプロセス管理の設定を取得する
func (c *Client) GetProcessManagement(ctx context.Context, params GetProcessManagementParams) (*GetProcessManagementResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/status", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetProcessManagementResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateProcessManagement はプロセス管理の設定を更新する（プレビュー環境）
func (c *Client) UpdateProcessManagement(ctx context.Context, params UpdateProcessManagementParams) (*UpdateProcessManagementResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Enable != nil {
		reqBody["enable"] = *params.Enable
	}
	if params.States != nil {
		reqBody["states"] = params.States
	}
	if params.Actions != nil {
		reqBody["actions"] = params.Actions
	}
	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/status", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateProcessManagementResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// --- 権限API ---

// GetAppAcl はアプリのアクセス権限を取得する
func (c *Client) GetAppAcl(ctx context.Context, params GetAppAclParams) (*GetAppAclResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	body, err := c.httpClient.GetWithBody(ctx, "app/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetAppAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateAppAcl はアプリのアクセス権限を更新する（プレビュー環境）
func (c *Client) UpdateAppAcl(ctx context.Context, params UpdateAppAclParams) (*UpdateAppAclResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"rights": params.Rights,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/app/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateAppAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetFieldAcl はフィールドのアクセス権限を取得する
func (c *Client) GetFieldAcl(ctx context.Context, params GetFieldAclParams) (*GetFieldAclResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	body, err := c.httpClient.GetWithBody(ctx, "field/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetFieldAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateFieldAcl はフィールドのアクセス権限を更新する（プレビュー環境）
func (c *Client) UpdateFieldAcl(ctx context.Context, params UpdateFieldAclParams) (*UpdateFieldAclResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"rights": params.Rights,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/field/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateFieldAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetRecordAcl はレコードのアクセス権限を取得する
func (c *Client) GetRecordAcl(ctx context.Context, params GetRecordAclParams) (*GetRecordAclResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody(ctx, "record/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetRecordAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateRecordAcl はレコードのアクセス権限を更新する（プレビュー環境）
func (c *Client) UpdateRecordAcl(ctx context.Context, params UpdateRecordAclParams) (*UpdateRecordAclResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"rights": params.Rights,
	}

	if params.Revision != "" {
		reqBody["revision"] = params.Revision
	}

	body, err := c.httpClient.Put(ctx, "preview/record/acl", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateRecordAclResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}
