// Package auth は認証方式を提供する
package auth

import (
	"encoding/base64"
	"net/http"
)

// Auth は認証インターフェース
type Auth interface {
	// Apply はリクエストに認証情報を付与する
	Apply(req *http.Request)
}

// APITokenAuth はAPIトークン認証
type APITokenAuth struct {
	Token string
}

// Apply はAPIトークンをヘッダーに設定する
func (a APITokenAuth) Apply(req *http.Request) {
	req.Header.Set("X-Cybozu-API-Token", a.Token)
}

// PasswordAuth はパスワード認証（kintone専用）
type PasswordAuth struct {
	Username string
	Password string
}

// Apply はBase64エンコードした認証情報をX-Cybozu-Authorizationヘッダーに設定する
func (a PasswordAuth) Apply(req *http.Request) {
	credentials := a.Username + ":" + a.Password
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Set("X-Cybozu-Authorization", encoded)
}

// BasicAuth はBasic認証（プロキシ等で使用）
type BasicAuth struct {
	Username string
	Password string
}

// Apply はBasic認証ヘッダーを設定する
func (a BasicAuth) Apply(req *http.Request) {
	credentials := a.Username + ":" + a.Password
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Set("Authorization", "Basic "+encoded)
}
