// 基本的な使用例
package main

import (
	"fmt"
	"os"

	"github.com/goqoo-on-kintone/goten"
	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/record"
	"github.com/joho/godotenv"
)

// MyRecord はサンプルのレコード構造体
type MyRecord struct {
	RecordNumber struct {
		Value string `json:"value"`
	} `json:"レコード番号"`
}

func main() {
	// .envファイルから環境変数を読み込む
	godotenv.Load()

	apiToken := os.Getenv("KINTONE_API_TOKEN")
	if apiToken == "" {
		fmt.Println("KINTONE_API_TOKENを設定してください")
		os.Exit(1)
	}

	// クライアント作成
	client := goten.NewClient(goten.Options{
		BaseURL: "https://example.cybozu.com",
		Auth:    auth.APITokenAuth{Token: apiToken},
	})

	// レコード取得（ジェネリクス使用）
	result, err := record.GetRecords[MyRecord](client.Record, record.GetRecordsParams{
		App:   "1",
		Query: "limit 10",
	})
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("取得したレコード数: %d\n", len(result.Records))
	for i, rec := range result.Records {
		fmt.Printf("レコード %d: %s\n", i+1, rec.RecordNumber.Value)
	}
}
