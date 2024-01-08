package todo

import (
	"fmt"
	"net/http"

	todoApp "github.com/KentaroKajiyama/Internship-go-api/application/todo"
	todoDi "github.com/KentaroKajiyama/Internship-go-api/di/todo"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/validator"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

type PostTodosParams struct {
	ID          string `query:"id"`
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

type PutTodosParams struct {
	ID          string `query:"id"`
	TodoID      string `query:"todo_id"`
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

type DeleteTodosParams struct {
	ID          string `query:"id"`
	TodoID      string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

// Post 新規作成
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が作成されない想定
func (h *TodoHandler) PostTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostTodosParams
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
	input_dto := todoApp.CreateTodoUseCaseInputDto{
		id:           params.ID,
		title:        params.Title,
		description:  params.Description,
		is_deletable: params.IsDeletable,
	}
	// UseCase処理
	err = todoDi.CreateTodo().Create(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？ Jsonデータを返した方がいいのか？
	return ctx.String(http.StatusOK, "Todo項目を新規作成しました。")
}

// PUT 更新
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が更新されない想定
func (h *TodoHandler) PutTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutTodosParams
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
	input_dto := todoApp.UpdateTodoUseCaseInputDto{
		id:           params.ID,
		todo_id:      params.TodoID,
		title:        params.Title,
		description:  params.Description,
		is_deletable: params.IsDeletable,
	}
	// UseCase処理
	err = todoDi.UpdateTodo().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？
	return ctx.String(http.StatusOK, "Todo項目を更新しました。")
}

// DELETE 削除
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtodo項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *TodoHandler) DeleteTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTodosParams
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
	input_dto := todoApp.DeleteTodoUseCaseInputDto{
		id:           params.ID,
		todo_id:      params.TodoID,
		is_deletable: params.IsDeletable,
	}
	// UseCase処理
	err = todoDi.DeleteTodo().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//レスポンスはなしでいいのか？
	return ctx.String(http.StatusOK, "Todo項目を削除しました。")
}
