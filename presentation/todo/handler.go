package todo

import (
	"fmt"
	"net/http"

	todoApp "github.com/KentaroKajiyama/Internship-go-api/application/todo"
	todoDi "github.com/KentaroKajiyama/Internship-go-api/di/todo"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

// Get Todo項目（１つ）の参照
func (h *TodoHandler) GetTodo(ctx echo.Context) error {
	//リクエストパラメーター取得（リクエストのボディに対するエラーハンドリング→データ型や形式等が合っているか？）
	var params GetTodoParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.FindTodoUseCaseInputDto{
		Id:     params.Id,
		TodoId: params.TodoId,
	}
	// UseCase処理
	todo, err := todoDi.FindTodo().Find(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// GET Todo項目（複数）の参照
func (h *TodoHandler) GetTodos(ctx echo.Context) error {
	//リクエストパラメーター取得（リクエストのボディに対するエラーハンドリング→データ型や形式等が合っているか？）
	var params GetTodosParams
	var trm TodosResponseModel
	var response []TodosResponseModel
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.FindTodosUseCaseInputDto{
		Id:     ctx.Get("id").(string),
		TodoId: params.TodoId,
		Title:  params.Title,
	}
	// UseCase処理
	todos, err := todoDi.FindTodos().FindMultiple(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	for _, todo := range todos {
		trm = TodosResponseModel{
			Id:          todo.Id(),
			TodoId:      todo.TodoId(),
			Title:       todo.Title(),
			Description: todo.Description(),
			IsDeletable: todo.IsDeletable(),
			CreatedAt:   todo.CreatedAt(),
			UpdatedAt:   todo.UpdatedAt(),
		}
		response = append(response, trm)
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// Post 新規作成
// 一度に一つしかtodo項目が作成されない想定
func (h *TodoHandler) PostTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostTodosParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.CreateTodoUseCaseInputDto{
		Id:          params.Id,
		Title:       params.Title,
		Description: params.Description,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理 ここでdbが挿入される
	todo, err := todoDi.CreateTodo().Create(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// PUT 更新
// 一度に一つしかtodo項目が更新されない想定
func (h *TodoHandler) PutTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutTodosParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.UpdateTodoUseCaseInputDto{
		Id:          params.Id,
		TodoId:      params.TodoId,
		Title:       params.Title,
		Description: params.Description,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理
	todo, err := todoDi.UpdateTodo().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// DELETE 削除
// 一度に一つしかtodo項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *TodoHandler) DeleteTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTodosParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.DeleteTodoUseCaseInputDto{
		Id:          params.Id,
		TodoId:      params.TodoId,
		IsDeletable: params.IsDeletable,
	}
	// UseCase処理
	todo, err := todoDi.DeleteTodo().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}
