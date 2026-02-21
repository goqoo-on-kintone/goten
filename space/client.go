// Package space はスペース管理APIを提供する
package space

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goqoo-on-kintone/goten/http"
)

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

// GetSpace はスペースの情報を取得する
func (c *Client) GetSpace(ctx context.Context, params GetSpaceParams) (*Space, error) {
	reqBody := map[string]any{
		"id": params.ID,
	}

	body, err := c.httpClient.GetWithBody(ctx, "space", reqBody)
	if err != nil {
		return nil, err
	}

	var result Space
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// GetSpaceMembers はスペースのメンバーを取得する
func (c *Client) GetSpaceMembers(ctx context.Context, params GetSpaceMembersParams) (*GetSpaceMembersResult, error) {
	reqBody := map[string]any{
		"id": params.ID,
	}

	body, err := c.httpClient.GetWithBody(ctx, "space/members", reqBody)
	if err != nil {
		return nil, err
	}

	var result GetSpaceMembersResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateSpace はスペースの設定を更新する
func (c *Client) UpdateSpace(ctx context.Context, params UpdateSpaceParams) error {
	reqBody := map[string]any{
		"id": params.ID,
	}

	if params.Name != nil {
		reqBody["name"] = *params.Name
	}
	if params.Body != nil {
		reqBody["body"] = *params.Body
	}
	if params.IsPrivate != nil {
		reqBody["isPrivate"] = *params.IsPrivate
	}
	if params.UseMultiThread != nil {
		reqBody["useMultiThread"] = *params.UseMultiThread
	}
	if params.FixedMember != nil {
		reqBody["fixedMember"] = *params.FixedMember
	}
	if params.ShowAnnouncement != nil {
		reqBody["showAnnouncement"] = *params.ShowAnnouncement
	}
	if params.ShowThreadList != nil {
		reqBody["showThreadList"] = *params.ShowThreadList
	}
	if params.ShowAppList != nil {
		reqBody["showAppList"] = *params.ShowAppList
	}
	if params.ShowMemberList != nil {
		reqBody["showMemberList"] = *params.ShowMemberList
	}
	if params.ShowRelatedLinkList != nil {
		reqBody["showRelatedLinkList"] = *params.ShowRelatedLinkList
	}

	_, err := c.httpClient.Put(ctx, "space", reqBody)
	return err
}

// UpdateSpaceMembers はスペースのメンバーを更新する
func (c *Client) UpdateSpaceMembers(ctx context.Context, params UpdateSpaceMembersParams) error {
	reqBody := map[string]any{
		"id":      params.ID,
		"members": params.Members,
	}

	_, err := c.httpClient.Put(ctx, "space/members", reqBody)
	return err
}

// AddThread はスレッドを追加する
func (c *Client) AddThread(ctx context.Context, params AddThreadParams) (*AddThreadResult, error) {
	reqBody := map[string]any{
		"space": params.Space,
		"name":  params.Name,
	}

	body, err := c.httpClient.Post(ctx, "space/thread", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddThreadResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}

// UpdateThread はスレッドを更新する
func (c *Client) UpdateThread(ctx context.Context, params UpdateThreadParams) error {
	reqBody := map[string]any{
		"id": params.ID,
	}

	if params.Name != nil {
		reqBody["name"] = *params.Name
	}
	if params.Body != nil {
		reqBody["body"] = *params.Body
	}

	_, err := c.httpClient.Put(ctx, "space/thread", reqBody)
	return err
}

// AddThreadComment はスレッドにコメントを追加する
func (c *Client) AddThreadComment(ctx context.Context, params AddThreadCommentParams) (*AddThreadCommentResult, error) {
	reqBody := map[string]any{
		"space":   params.Space,
		"thread":  params.Thread,
		"comment": params.Comment,
	}

	body, err := c.httpClient.Post(ctx, "space/thread/comment", reqBody)
	if err != nil {
		return nil, err
	}

	var result AddThreadCommentResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("レスポンス解析エラー: %w", err)
	}

	return &result, nil
}
