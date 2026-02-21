// Package record はレコード操作APIを提供する
package record

import "github.com/goqoo-on-kintone/goten/types"

// GetRecordParams はGetRecordのパラメータ
type GetRecordParams struct {
	App types.AppID
	ID  types.RecordID
}

// GetRecordsParams はGetRecordsのパラメータ
type GetRecordsParams struct {
	App        types.AppID
	Fields     []string
	Query      string
	TotalCount bool
}

// GetRecordsResult はGetRecordsの結果
type GetRecordsResult[T any] struct {
	Records    []T
	TotalCount *string
}

// GetAllRecordsParams はGetAllRecordsのパラメータ
type GetAllRecordsParams struct {
	App       types.AppID
	Fields    []string
	Condition string
	OrderBy   string
}

// AddRecordParams はAddRecordのパラメータ
type AddRecordParams struct {
	App    types.AppID
	Record map[string]types.FieldValue
}

// AddRecordResult はAddRecordの結果
type AddRecordResult struct {
	ID       string `json:"id"`
	Revision string `json:"revision"`
}

// AddRecordsParams はAddRecordsのパラメータ
type AddRecordsParams struct {
	App     types.AppID
	Records []map[string]types.FieldValue
}

// AddRecordsResult はAddRecordsの結果
type AddRecordsResult struct {
	IDs       []string `json:"ids"`
	Revisions []string `json:"revisions"`
}

// UpdateRecordParams はUpdateRecordのパラメータ
type UpdateRecordParams struct {
	App       types.AppID
	ID        types.RecordID
	UpdateKey *types.UpdateKey
	Record    map[string]types.FieldValue
	Revision  *types.Revision
}

// UpdateRecordResult はUpdateRecordの結果
type UpdateRecordResult struct {
	Revision string `json:"revision"`
}

// DeleteRecordsParams はDeleteRecordsのパラメータ
type DeleteRecordsParams struct {
	App       types.AppID
	IDs       []types.RecordID
	Revisions []types.Revision
}

// CreateCursorParams はCreateCursorのパラメータ
type CreateCursorParams struct {
	App    types.AppID
	Fields []string
	Query  string
	Size   int // 1回の取得件数（最大500、デフォルト100）
}

// CreateCursorResult はCreateCursorの結果
type CreateCursorResult struct {
	ID         string `json:"id"`
	TotalCount string `json:"totalCount"`
}

// GetRecordsByCursorParams はGetRecordsByCursorのパラメータ
type GetRecordsByCursorParams struct {
	ID string
}

// GetRecordsByCursorResult はGetRecordsByCursorの結果
type GetRecordsByCursorResult[T any] struct {
	Records []T  `json:"records"`
	Next    bool `json:"next"`
}

// DeleteCursorParams はDeleteCursorのパラメータ
type DeleteCursorParams struct {
	ID string
}
