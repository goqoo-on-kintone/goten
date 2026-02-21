package file_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/file"
	gotenhttp "github.com/goqoo-on-kintone/goten/http"
)

func TestUpload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("期待されるメソッド: POST, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/file.json" {
			t.Errorf("期待されるパス: /k/v1/file.json, 実際: %s", r.URL.Path)
		}

		contentType := r.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "multipart/form-data") {
			t.Errorf("期待されるContent-Type: multipart/form-data..., 実際: %s", contentType)
		}

		// ファイルを読み取る
		formFile, header, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("ファイル取得エラー: %v", err)
		}
		defer formFile.Close()

		if header.Filename != "test.txt" {
			t.Errorf("期待されるファイル名: test.txt, 実際: %s", header.Filename)
		}

		content, _ := io.ReadAll(formFile)
		if string(content) != "テストファイルの内容" {
			t.Errorf("期待される内容: テストファイルの内容, 実際: %s", string(content))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"fileKey": "test-file-key-12345"}`))
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := file.NewClient(httpClient)

	result, err := client.Upload(file.UploadParams{
		FileName: "test.txt",
		Reader:   strings.NewReader("テストファイルの内容"),
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if result.FileKey != "test-file-key-12345" {
		t.Errorf("期待されるFileKey: test-file-key-12345, 実際: %s", result.FileKey)
	}
}

func TestDownload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("期待されるメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.Path != "/k/v1/file.json" {
			t.Errorf("期待されるパス: /k/v1/file.json, 実際: %s", r.URL.Path)
		}

		fileKey := r.URL.Query().Get("fileKey")
		if fileKey != "test-file-key-12345" {
			t.Errorf("期待されるfileKey: test-file-key-12345, 実際: %s", fileKey)
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("ダウンロードされたファイルの内容"))
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := file.NewClient(httpClient)

	reader, err := client.Download(file.DownloadParams{
		FileKey: "test-file-key-12345",
	})

	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("読み取りエラー: %v", err)
	}

	if string(content) != "ダウンロードされたファイルの内容" {
		t.Errorf("期待される内容: ダウンロードされたファイルの内容, 実際: %s", string(content))
	}
}

func TestUploadError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": "CB_VA01", "id": "error-id", "message": "ファイルサイズが上限を超えています"}`))
	}))
	defer server.Close()

	httpClient := gotenhttp.NewDefaultClient(server.URL, auth.APITokenAuth{Token: "test-token"})
	client := file.NewClient(httpClient)

	_, err := client.Upload(file.UploadParams{
		FileName: "large.txt",
		Reader:   strings.NewReader("大きなファイル"),
	})

	if err == nil {
		t.Error("エラーが発生するはずが、発生しなかった")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "CB_VA01") {
		t.Errorf("エラーメッセージにコードが含まれていない: %s", errMsg)
	}
}
