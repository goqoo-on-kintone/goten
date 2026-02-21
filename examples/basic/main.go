// 基本的な使用例
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goqoo-on-kintone/goten"
	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/record"
	"github.com/joho/godotenv"
)

// InquiryRecord は問い合わせアプリのレコード構造体
type InquiryRecord struct {
	ID struct {
		Value string `json:"value"`
	} `json:"$id"`
	RecordNumber struct {
		Value string `json:"value"`
	} `json:"レコード番号"`
	InquiryType struct {
		Value string `json:"value"`
	} `json:"問い合わせ種別"`
	PersonName struct {
		Value string `json:"value"`
	} `json:"ご担当者名"`
	Status struct {
		Value string `json:"value"`
	} `json:"対応状況"`
	CreatedAt struct {
		Value string `json:"value"`
	} `json:"作成日時"`
}

func main() {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		// .envファイルがなくても続行
	}

	apiToken := os.Getenv("KINTONE_API_TOKEN")
	if apiToken == "" {
		fmt.Println("KINTONE_API_TOKENを設定してください")
		os.Exit(1)
	}

	// コンテキスト作成
	ctx := context.Background()

	// クライアント作成
	client := goten.NewClient(goten.Options{
		BaseURL: "https://the-red.cybozu.com",
		Auth:    auth.APITokenAuth{Token: apiToken},
	})

	// レコード取得（ジェネリクス使用）
	fmt.Println("=== レコード取得 ===")
	result, err := record.GetRecords[InquiryRecord](ctx, client.Record, record.GetRecordsParams{
		App:        "276",
		Query:      "limit 5",
		TotalCount: true,
	})
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		os.Exit(1)
	}

	if result.TotalCount != nil {
		fmt.Printf("総件数: %s\n", *result.TotalCount)
	}
	fmt.Printf("取得件数: %d\n\n", len(result.Records))

	for i, rec := range result.Records {
		fmt.Printf("--- レコード %d ---\n", i+1)
		fmt.Printf("  ID: %s\n", rec.ID.Value)
		fmt.Printf("  レコード番号: %s\n", rec.RecordNumber.Value)
		fmt.Printf("  問い合わせ種別: %s\n", rec.InquiryType.Value)
		fmt.Printf("  ご担当者名: %s\n", rec.PersonName.Value)
		fmt.Printf("  対応状況: %s\n", rec.Status.Value)
		fmt.Printf("  作成日時: %s\n", rec.CreatedAt.Value)
		fmt.Println()
	}

	// GetAllRecords（全件取得）のテスト
	fmt.Println("=== 全レコード取得 (GetAllRecords) ===")
	allRecords, err := record.GetAllRecords[InquiryRecord](ctx, client.Record, record.GetAllRecordsParams{
		App:     "276",
		OrderBy: "レコード番号 desc",
	})
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("全件数: %d\n", len(allRecords))

	// GetRecord（単一レコード取得）のテスト
	fmt.Println("\n=== 単一レコード取得 (GetRecord) ===")
	singleRecord, err := record.GetRecord[InquiryRecord](ctx, client.Record, record.GetRecordParams{
		App: "276",
		ID:  "1",
	})
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("レコードID: %s\n", singleRecord.ID.Value)
	fmt.Printf("ご担当者名: %s\n", singleRecord.PersonName.Value)
}
