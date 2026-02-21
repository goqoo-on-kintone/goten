// Package file はファイル操作APIを提供する
package file

import "github.com/goqoo-on-kintone/goten/http"

// Client はファイル操作クライアント
type Client struct {
	httpClient *http.DefaultClient
}

// NewClient は新しいFileClientを作成する
func NewClient(httpClient *http.DefaultClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// TODO: Upload, Download を実装
