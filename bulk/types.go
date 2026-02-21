// Package bulk はバルクリクエストAPIを提供する
package bulk

// Request はバルクリクエストの1つの要素
type Request struct {
	Method  string `json:"method"`  // GET, POST, PUT, DELETE
	API     string `json:"api"`     // /k/v1/record.json など
	Payload any    `json:"payload"` // リクエストボディ
}

// SendParams はSendのパラメータ
type SendParams struct {
	Requests []Request
}

// SendResult はSendの結果
type SendResult struct {
	Results []any `json:"results"`
}
