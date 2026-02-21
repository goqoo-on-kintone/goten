package http_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
)

func TestBuildPath(t *testing.T) {
	tests := []struct {
		name         string
		guestSpaceID *int
		endpoint     string
		wantPath     string
	}{
		{
			name:     "通常のパス",
			endpoint: "records",
			wantPath: "/k/v1/records.json",
		},
		{
			name:         "ゲストスペースのパス",
			guestSpaceID: intPtr(123),
			endpoint:     "records",
			wantPath:     "/k/guest/123/v1/records.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.wantPath {
					t.Errorf("期待されるパス: %s, 実際: %s", tt.wantPath, r.URL.Path)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{}"))
			}))
			defer server.Close()

			client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test"})
			if tt.guestSpaceID != nil {
				client.GuestSpaceID = tt.guestSpaceID
			}

			client.Get(tt.endpoint, nil)
		})
	}
}

func TestGetWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}

		// リクエストボディを確認
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("リクエストボディの解析エラー: %v", err)
		}

		if body["app"] != "1" {
			t.Errorf("期待されるapp: 1, 実際: %v", body["app"])
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result": "ok"}`))
	}))
	defer server.Close()

	client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test"})
	result, err := client.GetWithBody("records", map[string]any{"app": "1"})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}

	var response map[string]string
	json.Unmarshal(result, &response)
	if response["result"] != "ok" {
		t.Errorf("期待される結果: ok, 実際: %s", response["result"])
	}
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("期待されるContent-Type: application/json, 実際: %s", contentType)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id": "123"}`))
	}))
	defer server.Close()

	client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test"})
	result, err := client.Post("record", map[string]any{"app": "1"})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}

	var response map[string]string
	json.Unmarshal(result, &response)
	if response["id"] != "123" {
		t.Errorf("期待されるID: 123, 実際: %s", response["id"])
	}
}

func TestAuthHeaderIsSet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Cybozu-API-Token")
		if token != "my-secret-token" {
			t.Errorf("期待されるトークン: my-secret-token, 実際: %s", token)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "my-secret-token"})
	client.Get("records", nil)
}

func TestPostMultipart(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}

		contentType := r.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "multipart/form-data") {
			t.Errorf("期待されるContent-Type: multipart/form-data..., 実際: %s", contentType)
		}

		// ファイルを読み取る
		file, header, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("ファイル取得エラー: %v", err)
		}
		defer file.Close()

		if header.Filename != "test.txt" {
			t.Errorf("期待されるファイル名: test.txt, 実際: %s", header.Filename)
		}

		content, _ := io.ReadAll(file)
		if string(content) != "テストファイル内容" {
			t.Errorf("期待される内容: テストファイル内容, 実際: %s", string(content))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"fileKey": "abc123"}`))
	}))
	defer server.Close()

	client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test"})
	result, err := client.PostMultipart("file", "test.txt", strings.NewReader("テストファイル内容"))

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}

	var response map[string]string
	json.Unmarshal(result, &response)
	if response["fileKey"] != "abc123" {
		t.Errorf("期待されるfileKey: abc123, 実際: %s", response["fileKey"])
	}
}

func TestAPIErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": "CB_VA01", "id": "error-id", "message": "バリデーションエラー"}`))
	}))
	defer server.Close()

	client := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test"})
	_, err := client.Get("records", nil)

	if err == nil {
		t.Fatal("エラーが発生するはずが、発生しなかった")
	}

	// エラーメッセージを確認
	errMsg := err.Error()
	if !strings.Contains(errMsg, "CB_VA01") {
		t.Errorf("エラーメッセージにコードが含まれていない: %s", errMsg)
	}
}

// ヘルパー関数
func intPtr(i int) *int {
	return &i
}
