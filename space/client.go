// Package space はスペース管理APIを提供する
package space

import "github.com/goqoo-on-kintone/goten/http"

// Client はスペース管理クライアント
type Client struct {
	httpClient *http.DefaultClient
}

// NewClient は新しいSpaceClientを作成する
func NewClient(httpClient *http.DefaultClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// TODO: GetSpace, GetSpaceMembers などを実装
