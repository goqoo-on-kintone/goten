package bulk_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/bulk"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
)

func TestSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/bulkRequest.json" {
			t.Errorf("期待されるパス: /k/v1/bulkRequest.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		requests, ok := reqBody["requests"].([]any)
		if !ok || len(requests) != 2 {
			t.Errorf("期待されるリクエスト数: 2, 実際: %v", requests)
		}

		response := map[string]any{
			"results": []map[string]any{
				{"id": "1", "revision": "1"},
				{"id": "2", "revision": "1"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := bulk.NewClient(httpClient)

	result, err := client.Send(bulk.SendParams{
		Requests: []bulk.Request{
			{
				Method: "POST",
				API:    "/k/v1/record.json",
				Payload: map[string]any{
					"app":    "1",
					"record": map[string]any{},
				},
			},
			{
				Method: "POST",
				API:    "/k/v1/record.json",
				Payload: map[string]any{
					"app":    "1",
					"record": map[string]any{},
				},
			},
		},
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("期待される結果数: 2, 実際: %d", len(result.Results))
	}
}

func TestSendEmptyRequests(t *testing.T) {
	httpClient := gotenhttp.NewDefaultClient("http://localhost", auth.APITokenAuth{Token: "test-token"})
	client := bulk.NewClient(httpClient)

	_, err := client.Send(bulk.SendParams{
		Requests: []bulk.Request{},
	})

	if err == nil {
		t.Error("空のリクエストでエラーが発生するはずが、発生しなかった")
	}
}

func TestSendTooManyRequests(t *testing.T) {
	httpClient := gotenhttp.NewDefaultClient("http://localhost", auth.APITokenAuth{Token: "test-token"})
	client := bulk.NewClient(httpClient)

	// 21個のリクエストを作成（上限は20）
	requests := make([]bulk.Request, 21)
	for i := range requests {
		requests[i] = bulk.Request{
			Method:  "POST",
			API:     "/k/v1/record.json",
			Payload: map[string]any{},
		}
	}

	_, err := client.Send(bulk.SendParams{
		Requests: requests,
	})

	if err == nil {
		t.Error("上限超過でエラーが発生するはずが、発生しなかった")
	}
}

func TestBuilder(t *testing.T) {
	builder := bulk.NewBuilder()

	builder.
		AddRecord("1", map[string]any{"名前": map[string]any{"value": "テスト"}}).
		UpdateRecord("1", "100", map[string]any{"名前": map[string]any{"value": "更新"}}, "").
		DeleteRecords("1", []string{"101", "102"}, nil)

	requests := builder.Build()

	if len(requests) != 3 {
		t.Errorf("期待されるリクエスト数: 3, 実際: %d", len(requests))
	}

	// AddRecord
	if requests[0].Method != "POST" {
		t.Errorf("期待されるMethod: POST, 実際: %s", requests[0].Method)
	}
	if requests[0].API != "/k/v1/record.json" {
		t.Errorf("期待されるAPI: /k/v1/record.json, 実際: %s", requests[0].API)
	}

	// UpdateRecord
	if requests[1].Method != "PUT" {
		t.Errorf("期待されるMethod: PUT, 実際: %s", requests[1].Method)
	}

	// DeleteRecords
	if requests[2].Method != "DELETE" {
		t.Errorf("期待されるMethod: DELETE, 実際: %s", requests[2].Method)
	}
	if requests[2].API != "/k/v1/records.json" {
		t.Errorf("期待されるAPI: /k/v1/records.json, 実際: %s", requests[2].API)
	}
}

func TestBuilderCount(t *testing.T) {
	builder := bulk.NewBuilder()

	if builder.Count() != 0 {
		t.Errorf("期待されるCount: 0, 実際: %d", builder.Count())
	}

	builder.AddRecord("1", map[string]any{})
	builder.AddRecord("1", map[string]any{})

	if builder.Count() != 2 {
		t.Errorf("期待されるCount: 2, 実際: %d", builder.Count())
	}
}
