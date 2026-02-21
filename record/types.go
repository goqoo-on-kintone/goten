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

// UpdateRecordsParams はUpdateRecordsのパラメータ
type UpdateRecordsParams struct {
	App     types.AppID
	Records []UpdateRecordItem
}

// UpdateRecordItem は複数レコード更新時の各レコード
type UpdateRecordItem struct {
	ID        types.RecordID              `json:"id,omitempty"`
	UpdateKey *types.UpdateKey            `json:"updateKey,omitempty"`
	Record    map[string]types.FieldValue `json:"record"`
	Revision  *types.Revision             `json:"revision,omitempty"`
}

// UpdateRecordsResult はUpdateRecordsの結果
type UpdateRecordsResult struct {
	Records []UpdateRecordsResultItem `json:"records"`
}

// UpdateRecordsResultItem は更新結果の各レコード
type UpdateRecordsResultItem struct {
	ID       string `json:"id"`
	Revision string `json:"revision"`
}

// --- コメントAPI ---

// GetRecordCommentsParams はGetRecordCommentsのパラメータ
type GetRecordCommentsParams struct {
	App    types.AppID
	Record types.RecordID
	Order  string // asc または desc
	Offset int
	Limit  int // 最大10
}

// GetRecordCommentsResult はGetRecordCommentsの結果
type GetRecordCommentsResult struct {
	Comments []Comment `json:"comments"`
	Older    bool      `json:"older"`
	Newer    bool      `json:"newer"`
}

// Comment はコメント情報
type Comment struct {
	ID        string        `json:"id"`
	Text      string        `json:"text"`
	CreatedAt string        `json:"createdAt"`
	Creator   CommentUser   `json:"creator"`
	Mentions  []MentionUser `json:"mentions"`
}

// CommentUser はコメント作成者
type CommentUser struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// MentionUser はメンション対象
type MentionUser struct {
	Code string `json:"code"`
	Type string `json:"type"` // USER, GROUP, ORGANIZATION
}

// AddRecordCommentParams はAddRecordCommentのパラメータ
type AddRecordCommentParams struct {
	App      types.AppID
	Record   types.RecordID
	Comment  CommentContent
}

// CommentContent はコメント内容
type CommentContent struct {
	Text     string        `json:"text"`
	Mentions []MentionUser `json:"mentions,omitempty"`
}

// AddRecordCommentResult はAddRecordCommentの結果
type AddRecordCommentResult struct {
	ID string `json:"id"`
}

// DeleteRecordCommentParams はDeleteRecordCommentのパラメータ
type DeleteRecordCommentParams struct {
	App     types.AppID
	Record  types.RecordID
	Comment string // コメントID
}

// --- プロセス管理API ---

// UpdateRecordStatusParams はUpdateRecordStatusのパラメータ
type UpdateRecordStatusParams struct {
	App      types.AppID
	ID       types.RecordID
	Action   string          // アクション名
	Assignee string          // 次の作業者（省略可）
	Revision *types.Revision // リビジョン（省略可）
}

// UpdateRecordStatusResult はUpdateRecordStatusの結果
type UpdateRecordStatusResult struct {
	Revision string `json:"revision"`
}

// UpdateRecordsStatusParams はUpdateRecordsStatusのパラメータ
type UpdateRecordsStatusParams struct {
	App     types.AppID
	Records []UpdateStatusItem
}

// UpdateStatusItem はステータス更新対象のレコード
type UpdateStatusItem struct {
	ID       types.RecordID  `json:"id"`
	Action   string          `json:"action"`
	Assignee string          `json:"assignee,omitempty"`
	Revision *types.Revision `json:"revision,omitempty"`
}

// UpdateRecordsStatusResult はUpdateRecordsStatusの結果
type UpdateRecordsStatusResult struct {
	Records []UpdateRecordsResultItem `json:"records"`
}
