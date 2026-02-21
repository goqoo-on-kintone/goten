package space_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
	"github.com/goqoo-on-kintone/goten/space"
)

func TestGetSpace(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/space.json" {
			t.Errorf("期待されるパス: /k/v1/space.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["id"] != "1" {
			t.Errorf("期待されるid: 1, 実際: %v", reqBody["id"])
		}

		response := map[string]any{
			"id":             "1",
			"name":           "テストスペース",
			"defaultThread":  "100",
			"isPrivate":      false,
			"memberCount":    5,
			"useMultiThread": true,
			"isGuest":        false,
			"fixedMember":    false,
			"body":           "スペースの説明文",
			"creator": map[string]string{
				"code": "admin",
				"name": "管理者",
			},
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
	client := space.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetSpace(ctx, space.GetSpaceParams{
		ID: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.ID != "1" {
		t.Errorf("期待されるID: 1, 実際: %s", result.ID)
	}
	if result.Name != "テストスペース" {
		t.Errorf("期待されるName: テストスペース, 実際: %s", result.Name)
	}
	if result.MemberCount != 5 {
		t.Errorf("期待されるMemberCount: 5, 実際: %d", result.MemberCount)
	}
}

func TestGetSpaceMembers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/k/v1/space/members.json" {
			t.Errorf("期待されるパス: /k/v1/space/members.json, 実際: %s", r.URL.Path)
		}

		response := map[string]any{
			"members": []map[string]any{
				{
					"entity": map[string]string{
						"type": "USER",
						"code": "user1",
					},
					"isAdmin": true,
				},
				{
					"entity": map[string]string{
						"type": "USER",
						"code": "user2",
					},
					"isAdmin": false,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := space.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.GetSpaceMembers(ctx, space.GetSpaceMembersParams{
		ID: "1",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(result.Members) != 2 {
		t.Errorf("期待されるメンバー数: 2, 実際: %d", len(result.Members))
	}
	if result.Members[0].Entity.Code != "user1" {
		t.Errorf("期待されるCode: user1, 実際: %s", result.Members[0].Entity.Code)
	}
	if !result.Members[0].IsAdmin {
		t.Error("期待される値: true (isAdmin)")
	}
}

func TestUpdateSpace(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("期待されるメソッド: PUT, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/space.json" {
			t.Errorf("期待されるパス: /k/v1/space.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["id"] != "1" {
			t.Errorf("期待されるid: 1, 実際: %v", reqBody["id"])
		}
		if reqBody["name"] != "更新後のスペース名" {
			t.Errorf("期待されるname: 更新後のスペース名, 実際: %v", reqBody["name"])
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := space.NewClient(httpClient)

	ctx := context.Background()
	name := "更新後のスペース名"
	err := client.UpdateSpace(ctx, space.UpdateSpaceParams{
		ID:   "1",
		Name: &name,
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
}

func TestAddThread(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/space/thread.json" {
			t.Errorf("期待されるパス: /k/v1/space/thread.json, 実際: %s", r.URL.Path)
		}

		var reqBody map[string]any
		json.NewDecoder(r.Body).Decode(&reqBody)
		if reqBody["space"] != "1" {
			t.Errorf("期待されるspace: 1, 実際: %v", reqBody["space"])
		}
		if reqBody["name"] != "新しいスレッド" {
			t.Errorf("期待されるname: 新しいスレッド, 実際: %v", reqBody["name"])
		}

		response := map[string]string{
			"id": "999",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := space.NewClient(httpClient)

	ctx := context.Background()
	result, err := client.AddThread(ctx, space.AddThreadParams{
		Space: "1",
		Name:  "新しいスレッド",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.ID != "999" {
		t.Errorf("期待されるID: 999, 実際: %s", result.ID)
	}
}
