package todo

import (
	"fmt"
	"net/http"

	todoApp "github.com/KentaroKajiyama/Internship-go-api/application/todo"
	todoDi "github.com/KentaroKajiyama/Internship-go-api/di/todo"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/validator"
	"github.com/labstack/echo/v4"
)

type ToDoHandler struct {
}

func NewToDoHandler() *ToDoHandler {
	return &ToDoHandler{}
}

type PostToDosParams struct {
	ID          string `param:"id" query:"id" json:"id" form:"id" `
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

type PutToDosParams struct {
	ID          string `param:"id" query:"id" json:"id" form:"id"`
	ToDoID      string `param:"todo_id" query:"todo_id" json:"todo_id" form:"todo_id"`
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

type DeleteToDosParams struct {
	ID          string `param:"id" query:"id" json:"id" form:"id"`
	ToDoID      string `param:"todo_id" query:"todo_id" json:"todo_id" form:"todo_id"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

// Post 新規作成
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が作成されない想定
func (h *ToDoHandler) PostToDos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostToDosParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.CreateToDoUseCaseInputDto{
		ID:          params.ID,
		Title:       params.Title,
		Description: params.Description,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理 ここでdbが挿入される
	err = todoDi.CreateToDo().Create(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？ Jsonデータを返した方がいいのか？作成したデータの確認をするかどうかは
	return ctx.String(http.StatusOK, "Todo項目を新規作成しました。")
}

// PUT 更新
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が更新されない想定
func (h *ToDoHandler) PutToDos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutToDosParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//  Presentation -> UseCase
	input_dto := todoApp.UpdateToDoUseCaseInputDto{
		ID:          params.ID,
		TodoID:      params.ToDoID,
		Title:       params.Title,
		Description: params.Description,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理
	err = todoDi.UpdateToDo().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？
	return ctx.String(http.StatusOK, "Todo項目を更新しました。")
}

// DELETE 削除
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *ToDoHandler) DeleteToDos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteToDosParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//  Presentation -> UseCase
	input_dto := todoApp.DeleteToDoUseCaseInputDto{
		ID:          params.ID,
		TodoID:      params.ToDoID,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理
	err = todoDi.DeleteToDo().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？
	return ctx.String(http.StatusOK, "ToDo項目を削除しました。")
}
