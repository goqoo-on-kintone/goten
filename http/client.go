// Package http はHTTPクライアント抽象化を提供する
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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

// Get はGETリクエストを実行する（クエリパラメータ版）
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

// GetWithBody はGETリクエストを実行する（リクエストボディ版）
// kintone REST APIはGETでもリクエストボディを受け付ける
func (c *DefaultClient) GetWithBody(path string, body any) ([]byte, error) {
	url := c.buildPath(path)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

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

// Delete はDELETEリクエストを実行する（クエリパラメータ版）
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

// DeleteWithBody はDELETEリクエストを実行する（リクエストボディ版）
func (c *DefaultClient) DeleteWithBody(path string, body any) ([]byte, error) {
	url := c.buildPath(path)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	return c.do(req)
}

// PostMultipart はmultipart/form-dataでファイルをアップロードする
func (c *DefaultClient) PostMultipart(path string, fileName string, reader io.Reader) ([]byte, error) {
	url := c.buildPath(path)

	// multipartボディを作成
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("フォームファイル作成エラー: %w", err)
	}

	if _, err := io.Copy(part, reader); err != nil {
		return nil, fmt.Errorf("ファイルコピーエラー: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("マルチパートクローズエラー: %w", err)
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	c.Auth.Apply(req)
	req.Header.Set("Content-Type", writer.FormDataContentType())

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

// GetFile はファイルをダウンロードする
func (c *DefaultClient) GetFile(path string, fileKey string) (io.ReadCloser, error) {
	url := c.buildPath(path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	q := req.URL.Query()
	q.Add("fileKey", fileKey)
	req.URL.RawQuery = q.Encode()

	c.Auth.Apply(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエスト実行エラー: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var apiErr kintoneError.KintoneRestAPIError
		if err := json.Unmarshal(body, &apiErr); err == nil {
			apiErr.Status = resp.StatusCode
			return nil, &apiErr
		}
		return nil, fmt.Errorf("APIエラー (status=%d): %s", resp.StatusCode, string(body))
	}

	return resp.Body, nil
}
