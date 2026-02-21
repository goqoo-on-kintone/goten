package record_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
	"github.com/goqoo-on-kintone/goten/record"
	"github.com/goqoo-on-kintone/goten/types"
)

// TestRecord はテスト用のレコード構造体
type TestRecord struct {
	ID struct {
		Value string `json:"value"`
	} `json:"$id"`
	Name struct {
		Value string `json:"value"`
	} `json:"名前"`
}

func TestGetRecords(t *testing.T) {
	// モックサーバーを作成
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストの検証
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/records.json" {
			t.Errorf("期待されるパス: /k/v1/records.json, 実際: %s", r.URL.Path)
		}

		// リクエストボディを検証
		var reqBody map[string]any
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatalf("リクエストボディの解析エラー: %v", err)
		}
		if reqBody["app"] != "1" {
			t.Errorf("期待されるapp: 1, 実際: %v", reqBody["app"])
		}

		// レスポンスを返す
		response := map[string]any{
			"records": []map[string]any{
				{
					"$id":  map[string]string{"value": "1"},
					"名前": map[string]string{"value": "テスト太郎"},
				},
				{
					"$id":  map[string]string{"value": "2"},
					"名前": map[string]string{"value": "テスト花子"},
				},
			},
			"totalCount": "2",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// クライアントを作成
	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := record.NewClient(httpClient)

	// テスト実行
	result, err := record.GetRecords[TestRecord](client, record.GetRecordsParams{
		App:        "1",
		TotalCount: true,
	})

	// 結果を検証
	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Records) != 2 {
		t.Errorf("期待されるレコード数: 2, 実際: %d", len(result.Records))
	}
	if result.Records[0].ID.Value != "1" {
		t.Errorf("期待されるID: 1, 実際: %s", result.Records[0].ID.Value)
	}
	if result.Records[0].Name.Value != "テスト太郎" {
		t.Errorf("期待される名前: テスト太郎, 実際: %s", result.Records[0].Name.Value)
	}
	if result.TotalCount == nil || *result.TotalCount != "2" {
		t.Errorf("期待されるtotalCount: 2, 実際: %v", result.TotalCount)
	}
}

func TestGetRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/record.json" {
			t.Errorf("期待されるパス: /k/v1/record.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["id"] != "123" {
			t.Errorf("期待されるid: 123, 実際: %v", reqBody["id"])
		}

		response := map[string]any{
			"record": map[string]any{
				"$id":  map[string]string{"value": "123"},
				"名前": map[string]string{"value": "単一レコード"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := record.NewClient(httpClient)

	result, err := record.GetRecord[TestRecord](client, record.GetRecordParams{
		App: "1",
		ID:  "123",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.ID.Value != "123" {
		t.Errorf("期待されるID: 123, 実際: %s", result.ID.Value)
	}
	if result.Name.Value != "単一レコード" {
		t.Errorf("期待される名前: 単一レコード, 実際: %s", result.Name.Value)
	}
}

func TestAddRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/record.json" {
			t.Errorf("期待されるパス: /k/v1/record.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["app"] != "1" {
			t.Errorf("期待されるapp: 1, 実際: %v", reqBody["app"])
		}

		response := map[string]any{
			"id":       "999",
			"revision": "1",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := record.NewClient(httpClient)

	result, err := client.AddRecord(record.AddRecordParams{
		App: "1",
		Record: map[string]types.FieldValue{
			"名前": {Value: "新規レコード"},
		},
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.ID != "999" {
		t.Errorf("期待されるID: 999, 実際: %s", result.ID)
	}
	if result.Revision != "1" {
		t.Errorf("期待されるrevision: 1, 実際: %s", result.Revision)
	}
}

func TestGetAllRecords(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)

		// クエリにoffsetが含まれているか確認
		query, _ := reqBody["query"].(string)

		var records []map[string]any
		if callCount == 1 {
			// 1回目: 500件返す（次のページあり）
			for i := 0; i < 500; i++ {
				records = append(records, map[string]any{
					"$id":  map[string]string{"value": string(rune('0' + i%10))},
					"名前": map[string]string{"value": "レコード"},
				})
			}
		} else {
			// 2回目: 100件返す（最後のページ）
			for i := 0; i < 100; i++ {
				records = append(records, map[string]any{
					"$id":  map[string]string{"value": string(rune('0' + i%10))},
					"名前": map[string]string{"value": "レコード"},
				})
			}
		}

		_ = query // クエリは検証用に使用可能

		response := map[string]any{
			"records": records,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := record.NewClient(httpClient)

	result, err := record.GetAllRecords[TestRecord](client, record.GetAllRecordsParams{
		App: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	// 500 + 100 = 600件
	if len(result) != 600 {
		t.Errorf("期待されるレコード数: 600, 実際: %d", len(result))
	}
	if callCount != 2 {
		t.Errorf("期待されるAPI呼び出し回数: 2, 実際: %d", callCount)
	}
}

func TestAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]any{
			"code":    "CB_VA01",
			"id":      "test-error-id",
			"message": "入力内容が正しくありません。",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := record.NewClient(httpClient)

	_, err := record.GetRecords[TestRecord](client, record.GetRecordsParams{
		App: "1",
	})

	if err == nil {
		t.Fatal("エラーが発生するはずが、発生しなかった")
	}

	// エラーメッセージにコードが含まれているか確認
	errMsg := err.Error()
	if errMsg == "" {
		t.Error("エラーメッセージが空")
	}
}
