// Package app はアプリ設定APIを提供する
package app

import "github.com/goqoo-on-kintone/goten/http"

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

// TODO: GetFormFields, GetApp, GetViews などを実装
