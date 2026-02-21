// Package bulk はバルクリクエストAPIを提供する
package bulk

import "github.com/goqoo-on-kintone/goten/http"

// Client はバルクリクエストクライアント
type Client struct {
	httpClient *http.DefaultClient
}

// NewClient は新しいBulkRequestClientを作成する
func NewClient(httpClient *http.DefaultClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// TODO: Send を実装（最大20リクエスト）
