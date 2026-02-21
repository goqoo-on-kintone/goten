// Package bulk はバルクリクエストAPIを提供する
package bulk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goqoo-on-kintone/goten/http"
)

// MaxRequests はバルクリクエストの最大数
const MaxRequests = 20

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

// Send はバルクリクエストを実行する
// 最大20個のAPIリクエストを1回のリクエストで実行する
func (c *Client) Send(ctx context.Context, params SendParams) (*SendResult, error) {
	if len(params.Requests) == 0 {
		return nil, fmt.Errorf("リクエストが空です")
	}
	if len(params.Requests) > MaxRequests {
		return nil, fmt.Errorf("リクエスト数が上限(%d)を超えています: %d", MaxRequests, len(params.Requests))
	}

	reqBody := map[string]any{
		"requests": params.Requests,
	}

	body, err := c.httpClient.Post(ctx, "bulkRequest", reqBody)
	if err != nil {
		return nil, err
	}

	var result SendResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// Builder はバルクリクエストを構築するためのビルダー
type Builder struct {
	requests []Request
}

// NewBuilder は新しいBuilderを作成する
func NewBuilder() *Builder {
	return &Builder{
		requests: []Request{},
	}
}

// AddRecord はレコード追加リクエストを追加する
func (b *Builder) AddRecord(app string, record any) *Builder {
	b.requests = append(b.requests, Request{
		Method: "POST",
		API:    "/k/v1/record.json",
		Payload: map[string]any{
			"app":    app,
			"record": record,
		},
	})
	return b
}

// UpdateRecord はレコード更新リクエストを追加する
func (b *Builder) UpdateRecord(app string, id string, record any, revision string) *Builder {
	payload := map[string]any{
		"app":    app,
		"id":     id,
		"record": record,
	}
	if revision != "" {
		payload["revision"] = revision
	}
	b.requests = append(b.requests, Request{
		Method:  "PUT",
		API:     "/k/v1/record.json",
		Payload: payload,
	})
	return b
}

// DeleteRecords はレコード削除リクエストを追加する
func (b *Builder) DeleteRecords(app string, ids []string, revisions []string) *Builder {
	payload := map[string]any{
		"app": app,
		"ids": ids,
	}
	if len(revisions) > 0 {
		payload["revisions"] = revisions
	}
	b.requests = append(b.requests, Request{
		Method:  "DELETE",
		API:     "/k/v1/records.json",
		Payload: payload,
	})
	return b
}

// AddRequest はカスタムリクエストを追加する
func (b *Builder) AddRequest(method, api string, payload any) *Builder {
	b.requests = append(b.requests, Request{
		Method:  method,
		API:     api,
		Payload: payload,
	})
	return b
}

// Build はリクエスト一覧を返す
func (b *Builder) Build() []Request {
	return b.requests
}

// Count は現在のリクエスト数を返す
func (b *Builder) Count() int {
	return len(b.requests)
}
