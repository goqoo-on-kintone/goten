package auth_test

import (
	"net/http"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
)

func TestAPITokenAuth(t *testing.T) {
	a := auth.APITokenAuth{Token: "test-api-token"}

	req, _ := http.NewRequest("GET", "https://example.cybozu.com/k/v1/records.json", nil)
	a.Apply(req)

	// X-Cybozu-API-Tokenヘッダーが設定されているか確認
	token := req.Header.Get("X-Cybozu-API-Token")
	if token != "test-api-token" {
		t.Errorf("期待されるトークン: test-api-token, 実際: %s", token)
	}
}

func TestPasswordAuth(t *testing.T) {
	a := auth.PasswordAuth{
		Username: "testuser",
		Password: "changeme",
	}

	req, _ := http.NewRequest("GET", "https://example.cybozu.com/k/v1/records.json", nil)
	a.Apply(req)

	// X-Cybozu-Authorizationヘッダーが設定されているか確認
	authHeader := req.Header.Get("X-Cybozu-Authorization")
	if authHeader == "" {
		t.Error("X-Cybozu-Authorizationヘッダーが設定されていない")
	}

	// Base64エンコードされた値が含まれているか確認
	// "testuser:changeme" -> "dGVzdHVzZXI6Y2hhbmdlbWU="
	expected := "dGVzdHVzZXI6Y2hhbmdlbWU="
	if authHeader != expected {
		t.Errorf("期待される値: %s, 実際: %s", expected, authHeader)
	}
}

func TestBasicAuth(t *testing.T) {
	a := auth.BasicAuth{
		Username: "basicuser",
		Password: "basicpass",
	}

	req, _ := http.NewRequest("GET", "https://example.cybozu.com/k/v1/records.json", nil)
	a.Apply(req)

	// Authorizationヘッダーが設定されているか確認
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		t.Error("Authorizationヘッダーが設定されていない")
	}

	// "Basic " プレフィックスが含まれているか確認
	if len(authHeader) < 6 || authHeader[:6] != "Basic " {
		t.Errorf("Basicプレフィックスがない: %s", authHeader)
	}
}

func TestAuthInterface(t *testing.T) {
	// すべての認証方式がAuthインターフェースを実装しているか確認
	var _ auth.Auth = auth.APITokenAuth{}
	var _ auth.Auth = auth.PasswordAuth{}
	var _ auth.Auth = auth.BasicAuth{}
}
