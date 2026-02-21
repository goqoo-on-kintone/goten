// Package goten はkintone REST API向けGo言語SDKを提供する
package goten

import (
	"github.com/goqoo-on-kintone/goten/app"
	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/bulk"
	"github.com/goqoo-on-kintone/goten/file"
	"github.com/goqoo-on-kintone/goten/http"
	"github.com/goqoo-on-kintone/goten/record"
	"github.com/goqoo-on-kintone/goten/space"
)

// Client はkintone REST APIクライアント（ファサード）
type Client struct {
	Record *record.Client
	App    *app.Client
	Space  *space.Client
	File   *file.Client
	Bulk   *bulk.Client

	httpClient *http.DefaultClient
}

// Options はクライアント作成オプション
type Options struct {
	BaseURL      string
	Auth         auth.Auth
	GuestSpaceID *int
}

// NewClient は新しいClientを作成する
func NewClient(opts Options) *Client {
	httpClient := http.NewDefaultClient(opts.BaseURL, opts.Auth)
	if opts.GuestSpaceID != nil {
		httpClient.GuestSpaceID = opts.GuestSpaceID
	}

	return &Client{
		Record:     record.NewClient(httpClient),
		App:        app.NewClient(httpClient),
		Space:      space.NewClient(httpClient),
		File:       file.NewClient(httpClient),
		Bulk:       bulk.NewClient(httpClient),
		httpClient: httpClient,
	}
}
