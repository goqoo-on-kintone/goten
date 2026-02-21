package record

import (
	"encoding/json"
	"fmt"

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
	// kintone REST APIはGETでもリクエストボディを使用可能
	reqBody := map[string]any{
		"app": params.App,
	}

	if len(params.Fields) > 0 {
		reqBody["fields"] = params.Fields
	}
	if params.Query != "" {
		reqBody["query"] = params.Query
	}
	if params.TotalCount {
		reqBody["totalCount"] = true
	}

	body, err := c.httpClient.GetWithBody("records", reqBody)
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

	reqBody := map[string]any{
		"app": params.App,
		"id":  params.ID,
	}

	body, err := c.httpClient.GetWithBody("record", reqBody)
	if err != nil {
		return zero, err
	}

	var response getRecordResponse[T]
	if err := json.Unmarshal(body, &response); err != nil {
		return zero, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return response.Record, nil
}

// GetAllRecords は全レコードを取得する（ページング自動処理）
// 内部的に500件ずつ取得してすべて結合する
func GetAllRecords[T any](c *Client, params GetAllRecordsParams) ([]T, error) {
	const limit = 500
	var allRecords []T
	offset := 0

	for {
		// クエリを構築（offset/limitを追加）
		query := params.Condition
		if params.OrderBy != "" {
			if query != "" {
				query += " "
			}
			query += "order by " + params.OrderBy
		}
		if query != "" {
			query += " "
		}
		query += fmt.Sprintf("limit %d offset %d", limit, offset)

		result, err := GetRecords[T](c, GetRecordsParams{
			App:    params.App,
			Fields: params.Fields,
			Query:  query,
		})
		if err != nil {
			return nil, err
		}

		allRecords = append(allRecords, result.Records...)

		// 取得件数が上限未満なら終了
		if len(result.Records) < limit {
			break
		}

		offset += limit
	}

	return allRecords, nil
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

	_, err := c.httpClient.DeleteWithBody("records", reqBody)
	return err
}

// CreateCursor はカーソルを作成する
func (c *Client) CreateCursor(params CreateCursorParams) (*CreateCursorResult, error) {
	reqBody := map[string]any{
		"app": params.App,
	}

	if len(params.Fields) > 0 {
		reqBody["fields"] = params.Fields
	}
	if params.Query != "" {
		reqBody["query"] = params.Query
	}
	if params.Size > 0 {
		reqBody["size"] = params.Size
	}

	body, err := c.httpClient.Post("records/cursor", reqBody)
	if err != nil {
		return nil, err
	}

	var result CreateCursorResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetRecordsByCursor はカーソルを使ってレコードを取得する
func GetRecordsByCursor[T any](c *Client, params GetRecordsByCursorParams) (*GetRecordsByCursorResult[T], error) {
	reqBody := map[string]any{
		"id": params.ID,
	}

	body, err := c.httpClient.GetWithBody("records/cursor", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetRecordsByCursorResult[T]
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// DeleteCursor はカーソルを削除する
func (c *Client) DeleteCursor(params DeleteCursorParams) error {
	reqBody := map[string]any{
		"id": params.ID,
	}

	_, err := c.httpClient.DeleteWithBody("records/cursor", reqBody)
	return err
}
