// Package auth は認証方式を提供する
package auth

import "net/http"

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

// PasswordAuth はパスワード認証
type PasswordAuth struct {
	Username string
	Password string
}

// Apply はBase64エンコードした認証情報をヘッダーに設定する
func (a PasswordAuth) Apply(req *http.Request) {
	// TODO: Base64エンコード実装
	// X-Cybozu-Authorization: Base64(username:password)
}
