package record

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goqoo-on-kintone/goten/http"
)

// Client はレコード操作クライアント
type Client struct {
	httpClient *http.DefaultClient
}

// NewClient は新しいRecordClientを作成する
func NewClient(httpClient *http.DefaultClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// getRecordsResponse はレコード取得APIのレスポンス
type getRecordsResponse[T any] struct {
	Records    []T     `json:"records"`
	TotalCount *string `json:"totalCount"`
}

// GetRecords は複数レコードを取得する（ジェネリクス版）
func GetRecords[T any](c *Client, params GetRecordsParams) (*GetRecordsResult[T], error) {
	queryParams := map[string]string{
		"app": params.App,
	}

	if len(params.Fields) > 0 {
		queryParams["fields"] = strings.Join(params.Fields, ",")
	}
	if params.Query != "" {
		queryParams["query"] = params.Query
	}
	if params.TotalCount {
		queryParams["totalCount"] = "true"
	}

	body, err := c.httpClient.Get("records", queryParams)
	if err != nil {
		return nil, err
	}

	var response getRecordsResponse[T]
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &GetRecordsResult[T]{
		Records:    response.Records,
		TotalCount: response.TotalCount,
	}, nil
}

// getRecordResponse は単一レコード取得APIのレスポンス
type getRecordResponse[T any] struct {
	Record T `json:"record"`
}

// GetRecord は単一レコードを取得する（ジェネリクス版）
func GetRecord[T any](c *Client, params GetRecordParams) (T, error) {
	var zero T

	queryParams := map[string]string{
		"app": params.App,
		"id":  params.ID,
	}

	body, err := c.httpClient.Get("record", queryParams)
	if err != nil {
		return zero, err
	}

	var response getRecordResponse[T]
	if err := json.Unmarshal(body, &response); err != nil {
		return zero, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return response.Record, nil
}

// AddRecord はレコードを1件追加する
func (c *Client) AddRecord(params AddRecordParams) (*AddRecordResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"record": params.Record,
	}

	body, err := c.httpClient.Post("record", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddRecordResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// AddRecords はレコードを複数追加する（最大100件）
func (c *Client) AddRecords(params AddRecordsParams) (*AddRecordsResult, error) {
	reqBody := map[string]any{
		"app":     params.App,
		"records": params.Records,
	}

	body, err := c.httpClient.Post("records", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddRecordsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateRecord はレコードを1件更新する
func (c *Client) UpdateRecord(params UpdateRecordParams) (*UpdateRecordResult, error) {
	reqBody := map[string]any{
		"app":    params.App,
		"record": params.Record,
	}

	if params.ID != "" {
		reqBody["id"] = params.ID
	}
	if params.UpdateKey != nil {
		reqBody["updateKey"] = params.UpdateKey
	}
	if params.Revision != nil {
		reqBody["revision"] = *params.Revision
	}

	body, err := c.httpClient.Put("record", reqBody)
	if err != nil {
		return nil, err
	}

	var result UpdateRecordResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// DeleteRecords はレコードを削除する（最大100件）
func (c *Client) DeleteRecords(params DeleteRecordsParams) error {
	reqBody := map[string]any{
		"app": params.App,
		"ids": params.IDs,
	}

	if len(params.Revisions) > 0 {
		reqBody["revisions"] = params.Revisions
	}

	_, err := c.httpClient.Delete("records", nil)
	return err
}
