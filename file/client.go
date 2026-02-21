// Package file はファイル操作APIを提供する
package file

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/goqoo-on-kintone/goten/http"
)

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

// UploadParams はUploadのパラメータ
type UploadParams struct {
	FileName string
	Reader   io.Reader
}

// UploadResult はUploadの結果
type UploadResult struct {
	FileKey string `json:"fileKey"`
}

// Upload はファイルをアップロードする
func (c *Client) Upload(params UploadParams) (*UploadResult, error) {
	body, err := c.httpClient.PostMultipart("file", params.FileName, params.Reader)
	if err != nil {
		return nil, err
	}

	var result UploadResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// DownloadParams はDownloadのパラメータ
type DownloadParams struct {
	FileKey string
}

// Download はファイルをダウンロードする
func (c *Client) Download(params DownloadParams) (io.ReadCloser, error) {
	return c.httpClient.GetFile("file", params.FileKey)
}
