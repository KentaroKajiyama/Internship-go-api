package todo

import "time"

type TodosResponseModel struct {
	Id          string    `json:"id"`
	TodoId      string    `json:"todo_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDeletable bool      `json:"is_deletable"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 複数個レスポンスしたい場合は？ 作成時間などは返す必要なし？
