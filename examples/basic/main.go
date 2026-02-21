// 基本的な使用例
package main

import (
	"fmt"
	"os"

	"github.com/goqoo-on-kintone/goten"
	"github.com/goqoo-on-kintone/goten/auth"
	"github.com/goqoo-on-kintone/goten/record"
	"github.com/goqoo-on-kintone/gotenks/types"
	"github.com/joho/godotenv"
)

// InquiryRecord は問い合わせアプリのレコード構造体
type InquiryRecord struct {
	ID           types.IDField             `json:"$id"`
	Revision     types.RevisionField       `json:"$revision"`
	RecordNumber types.RecordNumberField   `json:"レコード番号"`
	InquiryType  types.RadioButtonField    `json:"問い合わせ種別"`
	PersonName   types.SingleLineTextField `json:"ご担当者名"`
	Detail       types.MultiLineTextField  `json:"詳細"`
	Status       types.DropDownField       `json:"対応状況"`
	ReceivedAt   types.DateTimeField       `json:"受付日時"`
	CreatedAt    types.CreatedTimeField    `json:"作成日時"`
	UpdatedAt    types.UpdatedTimeField    `json:"更新日時"`
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

	// クライアント作成
	client := goten.NewClient(goten.Options{
		BaseURL: "https://the-red.cybozu.com",
		Auth:    auth.APITokenAuth{Token: apiToken},
	})

	// レコード取得（ジェネリクス使用）
	fmt.Println("=== レコード取得 ===")
	result, err := record.GetRecords[InquiryRecord](client.Record, record.GetRecordsParams{
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
	allRecords, err := record.GetAllRecords[InquiryRecord](client.Record, record.GetAllRecordsParams{
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
	singleRecord, err := record.GetRecord[InquiryRecord](client.Record, record.GetRecordParams{
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
