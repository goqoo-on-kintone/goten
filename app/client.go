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
