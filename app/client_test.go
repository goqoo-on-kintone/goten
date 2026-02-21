package app_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goqoo-on-kintone/goten/app"
	"github.com/goqoo-on-kintone/goten/auth"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
)

func TestGetApp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/app.json" {
			t.Errorf("期待されるパス: /k/v1/app.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["id"] != "1" {
			t.Errorf("期待されるid: 1, 実際: %v", reqBody["id"])
		}

		response := map[string]any{
			"appId":       "1",
			"code":        "TEST_APP",
			"name":        "テストアプリ",
			"description": "テスト用のアプリです",
			"createdAt":   "2024-01-01T00:00:00Z",
			"creator": map[string]string{
				"code": "admin",
				"name": "管理者",
			},
			"modifiedAt": "2024-01-02T00:00:00Z",
			"modifier": map[string]string{
				"code": "admin",
				"name": "管理者",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := app.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetApp(ctx, app.GetAppParams{
		App: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.AppID != "1" {
		t.Errorf("期待されるAppID: 1, 実際: %s", result.AppID)
	}
	if result.Code != "TEST_APP" {
		t.Errorf("期待されるCode: TEST_APP, 実際: %s", result.Code)
	}
	if result.Name != "テストアプリ" {
		t.Errorf("期待されるName: テストアプリ, 実際: %s", result.Name)
	}
}

func TestGetApps(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/apps.json" {
			t.Errorf("期待されるパス: /k/v1/apps.json, 実際: %s", r.URL.Path)
		}

		response := map[string]any{
			"apps": []map[string]any{
				{
					"appId":       "1",
					"code":        "APP1",
					"name":        "アプリ1",
					"description": "",
					"createdAt":   "2024-01-01T00:00:00Z",
					"creator":     map[string]string{"code": "admin", "name": "管理者"},
					"modifiedAt":  "2024-01-02T00:00:00Z",
					"modifier":    map[string]string{"code": "admin", "name": "管理者"},
				},
				{
					"appId":       "2",
					"code":        "APP2",
					"name":        "アプリ2",
					"description": "",
					"createdAt":   "2024-01-01T00:00:00Z",
					"creator":     map[string]string{"code": "admin", "name": "管理者"},
					"modifiedAt":  "2024-01-02T00:00:00Z",
					"modifier":    map[string]string{"code": "admin", "name": "管理者"},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := app.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetApps(ctx, app.GetAppsParams{
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Apps) != 2 {
		t.Errorf("期待されるアプリ数: 2, 実際: %d", len(result.Apps))
	}
	if result.Apps[0].AppID != "1" {
		t.Errorf("期待されるAppID: 1, 実際: %s", result.Apps[0].AppID)
	}
}

func TestGetFormFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/k/v1/app/form/fields.json" {
			t.Errorf("期待されるパス: /k/v1/app/form/fields.json, 実際: %s", r.URL.Path)
		}

		response := map[string]any{
			"properties": map[string]any{
				"文字列": map[string]any{
					"type":     "SINGLE_LINE_TEXT",
					"code":     "文字列",
					"label":    "文字列フィールド",
					"required": true,
				},
				"数値": map[string]any{
					"type":  "NUMBER",
					"code":  "数値",
					"label": "数値フィールド",
				},
			},
			"revision": "5",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := app.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetFormFields(ctx, app.GetFormFieldsParams{
		App: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Properties) != 2 {
		t.Errorf("期待されるフィールド数: 2, 実際: %d", len(result.Properties))
	}
	if result.Properties["文字列"].Type != "SINGLE_LINE_TEXT" {
		t.Errorf("期待されるType: SINGLE_LINE_TEXT, 実際: %s", result.Properties["文字列"].Type)
	}
	if result.Revision != "5" {
		t.Errorf("期待されるRevision: 5, 実際: %s", result.Revision)
	}
}

func TestGetViews(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/k/v1/app/views.json" {
			t.Errorf("期待されるパス: /k/v1/app/views.json, 実際: %s", r.URL.Path)
		}

		response := map[string]any{
			"views": map[string]any{
				"(すべて)": map[string]any{
					"id":          "12345",
					"type":        "LIST",
					"name":        "(すべて)",
					"builtinType": "ASSIGNEE",
					"fields":      []string{"文字列", "数値"},
					"filterCond":  "",
					"sort":        "レコード番号 desc",
					"index":       "0",
				},
			},
			"revision": "10",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := app.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetViews(ctx, app.GetViewsParams{
		App: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Views) != 1 {
		t.Errorf("期待される一覧数: 1, 実際: %d", len(result.Views))
	}
	if result.Views["(すべて)"].Type != "LIST" {
		t.Errorf("期待されるType: LIST, 実際: %s", result.Views["(すべて)"].Type)
	}
	if result.Revision != "10" {
		t.Errorf("期待されるRevision: 10, 実際: %s", result.Revision)
	}
}
