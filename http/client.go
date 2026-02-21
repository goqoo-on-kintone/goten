// Package http はHTTPクライアント抽象化を提供する
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/goqoo-on-kintone/goten/auth"
	kintoneError "github.com/goqoo-on-kintone/goten/error"
)

// Client はHTTPクライアントインターフェース
type Client interface {
	Get(path string, params map[string]string) ([]byte, error)
	Post(path string, body any) ([]byte, error)
	Put(path string, body any) ([]byte, error)
	Delete(path string, params map[string]string) ([]byte, error)
}

// DefaultClient はデフォルトのHTTPクライアント実装
type DefaultClient struct {
	BaseURL      string
	Auth         auth.Auth
	GuestSpaceID *int
	HTTPClient   *http.Client
}

// NewDefaultClient は新しいDefaultClientを作成する
func NewDefaultClient(baseURL string, a auth.Auth) *DefaultClient {
	return &DefaultClient{
		BaseURL:    baseURL,
		Auth:       a,
		HTTPClient: &http.Client{},
	}
}

// buildPath はAPIパスを構築する
func (c *DefaultClient) buildPath(endpointName string) string {
	if c.GuestSpaceID != nil {
		return fmt.Sprintf("%s/k/guest/%d/v1/%s.json", c.BaseURL, *c.GuestSpaceID, endpointName)
	}
	return fmt.Sprintf("%s/k/v1/%s.json", c.BaseURL, endpointName)
}

// do はHTTPリクエストを実行する
func (c *DefaultClient) do(req *http.Request) ([]byte, error) {
	c.Auth.Apply(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエスト実行エラー: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンス読み取りエラー: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr kintoneError.KintoneRestAPIError
		if err := json.Unmarshal(body, &apiErr); err == nil {
			apiErr.Status = resp.StatusCode
			return nil, &apiErr
		}
		return nil, fmt.Errorf("APIエラー (status=%d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// Get はGETリクエストを実行する
func (c *DefaultClient) Get(path string, params map[string]string) ([]byte, error) {
	url := c.buildPath(path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return c.do(req)
}

// Post はPOSTリクエストを実行する
func (c *DefaultClient) Post(path string, body any) ([]byte, error) {
	url := c.buildPath(path)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	return c.do(req)
}

// Put はPUTリクエストを実行する
func (c *DefaultClient) Put(path string, body any) ([]byte, error) {
	url := c.buildPath(path)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	return c.do(req)
}

// Delete はDELETEリクエストを実行する
func (c *DefaultClient) Delete(path string, params map[string]string) ([]byte, error) {
	url := c.buildPath(path)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return c.do(req)
}
