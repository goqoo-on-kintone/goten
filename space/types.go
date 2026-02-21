// Package space はスペース管理APIを提供する
package space

// GetSpaceParams はGetSpaceのパラメータ
type GetSpaceParams struct {
	ID string
}

// Space はスペースの情報
type Space struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	DefaultThread       string `json:"defaultThread"`
	IsPrivate           bool   `json:"isPrivate"`
	Creator             User   `json:"creator"`
	Modifier            User   `json:"modifier"`
	MemberCount         int    `json:"memberCount"`
	CoverType           string `json:"coverType"`
	CoverKey            string `json:"coverKey,omitempty"`
	CoverURL            string `json:"coverUrl,omitempty"`
	Body                string `json:"body"`
	UseMultiThread      bool   `json:"useMultiThread"`
	IsGuest             bool   `json:"isGuest"`
	FixedMember         bool   `json:"fixedMember"`
	ShowAnnouncement    bool   `json:"showAnnouncement,omitempty"`
	ShowThreadList      bool   `json:"showThreadList,omitempty"`
	ShowAppList         bool   `json:"showAppList,omitempty"`
	ShowMemberList      bool   `json:"showMemberList,omitempty"`
	ShowRelatedLinkList bool   `json:"showRelatedLinkList,omitempty"`
}

// User はユーザー情報
type User struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetSpaceMembersParams はGetSpaceMembersのパラメータ
type GetSpaceMembersParams struct {
	ID string
}

// SpaceMember はスペースメンバー
type SpaceMember struct {
	Entity      Entity `json:"entity"`
	IsAdmin     bool   `json:"isAdmin"`
	IncludeSubs bool   `json:"includeSubs,omitempty"`
	IsImplicit  bool   `json:"isImplicit,omitempty"`
}

// Entity はエンティティ（ユーザー、グループ、組織）
type Entity struct {
	Type string `json:"type"` // USER, GROUP, ORGANIZATION
	Code string `json:"code"`
}

// GetSpaceMembersResult はGetSpaceMembersの結果
type GetSpaceMembersResult struct {
	Members []SpaceMember `json:"members"`
}

// UpdateSpaceParams はUpdateSpaceのパラメータ
type UpdateSpaceParams struct {
	ID                  string
	Name                *string `json:"name,omitempty"`
	Body                *string `json:"body,omitempty"`
	IsPrivate           *bool   `json:"isPrivate,omitempty"`
	UseMultiThread      *bool   `json:"useMultiThread,omitempty"`
	FixedMember         *bool   `json:"fixedMember,omitempty"`
	ShowAnnouncement    *bool   `json:"showAnnouncement,omitempty"`
	ShowThreadList      *bool   `json:"showThreadList,omitempty"`
	ShowAppList         *bool   `json:"showAppList,omitempty"`
	ShowMemberList      *bool   `json:"showMemberList,omitempty"`
	ShowRelatedLinkList *bool   `json:"showRelatedLinkList,omitempty"`
}

// UpdateSpaceMembersParams はUpdateSpaceMembersのパラメータ
type UpdateSpaceMembersParams struct {
	ID      string
	Members []SpaceMemberForUpdate
}

// SpaceMemberForUpdate は更新用のスペースメンバー
type SpaceMemberForUpdate struct {
	Entity      Entity `json:"entity"`
	IsAdmin     bool   `json:"isAdmin,omitempty"`
	IncludeSubs bool   `json:"includeSubs,omitempty"`
}

// AddThreadParams はAddThreadのパラメータ
type AddThreadParams struct {
	Space string
	Name  string
}

// AddThreadResult はAddThreadの結果
type AddThreadResult struct {
	ID string `json:"id"`
}

// UpdateThreadParams はUpdateThreadのパラメータ
type UpdateThreadParams struct {
	ID   string
	Name *string
	Body *string
}

// AddThreadCommentParams はAddThreadCommentのパラメータ
type AddThreadCommentParams struct {
	Space   string
	Thread  string
	Comment ThreadComment
}

// ThreadComment はスレッドコメント
type ThreadComment struct {
	Text     string    `json:"text"`
	Mentions []Mention `json:"mentions,omitempty"`
	Files    []File    `json:"files,omitempty"`
}

// Mention はメンション
type Mention struct {
	Code string `json:"code"`
	Type string `json:"type"` // USER, GROUP, ORGANIZATION
}

// File はファイル情報
type File struct {
	FileKey string `json:"fileKey"`
	Width   int    `json:"width,omitempty"`
}

// AddThreadCommentResult はAddThreadCommentの結果
type AddThreadCommentResult struct {
	ID string `json:"id"`
}
