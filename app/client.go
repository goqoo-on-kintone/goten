// Package app はアプリ設定APIを提供する
package app

import (
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
func (c *Client) GetApp(params GetAppParams) (*App, error) {
	reqBody := map[string]any{
		"id": params.App,
	}

	body, err := c.httpClient.GetWithBody("app", reqBody)
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
func (c *Client) GetApps(params GetAppsParams) (*GetAppsResult, error) {
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

	body, err := c.httpClient.GetWithBody("apps", reqBody)
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
func (c *Client) GetFormFields(params GetFormFieldsParams) (*GetFormFieldsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody("app/form/fields", reqBody)
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
func (c *Client) GetFormLayout(params GetFormLayoutParams) (*GetFormLayoutResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	body, err := c.httpClient.GetWithBody("app/form/layout", reqBody)
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
func (c *Client) GetViews(params GetViewsParams) (*GetViewsResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if params.Lang != "" {
		reqBody["lang"] = params.Lang
	}

	body, err := c.httpClient.GetWithBody("app/views", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetViewsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}
