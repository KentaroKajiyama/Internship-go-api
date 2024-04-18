package todo

import (
	"fmt"
	"net/http"
	"strconv"

	todoApp "github.com/KentaroKajiyama/Internship-go-api/application/todo"
	todoDi "github.com/KentaroKajiyama/Internship-go-api/di/todo"
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

// Get Todo項目（１つ）の参照
func (h *TodoHandler) GetTodo(ctx echo.Context) error {
	//Request
	var params GetTodoParams
	var responseTagIds []string
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Validation
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.FindTodoUseCaseInputDto{
		Id:     params.Id,
		TodoId: params.TodoId,
	}
	input_dto_for_get_tags := todoApp.GetTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
	}
	// Usecase
	todo, err := todoDi.FindTodo().Find(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	tags, err := todoDi.GetTagsInTodo().GetTagsInTodo(ctx.Request().Context(), input_dto_for_get_tags)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		IsChecked:   false,
		TagIds:      responseTagIds,
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) GetTagsInTodo(ctx echo.Context) error {
	var params GetTagsInTodoParams
	var responseTagIds []string
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	input_dto := todoApp.GetTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
	}
	tags, err := todoDi.GetTagsInTodo().GetTagsInTodo(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	response := TagsInTodoResponseModel{
		TodoId: tags.TodoId(),
		TagIds: responseTagIds,
	}
	return ctx.JSON(http.StatusOK, response)
}

// POST Add Tags to a specific todo
func (h *TodoHandler) PostTagsInTodo(ctx echo.Context) error {
	var params PostTagsInTodoParams
	var responseTagIds []string
	// Binding the parameters
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validation
	if err := ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Presentation -> UseCase
	tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.TagIds {
		tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		tagids_for_use_case = append(tagids_for_use_case, tagid_for_use_case)
	}
	input_dto := todoApp.AddTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
		TagIds: tagids_for_use_case,
	}
	// UseCase
	tags, err := todoDi.AddTagsInTodo().AddTagsInTodo(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase -> Presentation
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	response := TagsInTodoResponseModel{
		TodoId: tags.TodoId(),
		TagIds: responseTagIds,
	}
	// Response in JSON format
	return ctx.JSON(http.StatusOK, response)
}

// DELETE Delete Tags from a specific todo
func (h *TodoHandler) DeleteTagsInTodo(ctx echo.Context) error {
	var params DeleteTagsInTodoParams
	var responseTagIds []string
	// Binding the parameters
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validation
	if err := ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Presentation -> UseCase
	tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.TagIds {
		tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		tagids_for_use_case = append(tagids_for_use_case, tagid_for_use_case)
	}
	input_dto := todoApp.DeleteTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
		TagIds: tagids_for_use_case,
	}
	// UseCase
	tags, err := todoDi.DeleteTagsInTodo().DeleteTagsInTodo(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase -> Presentation
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	response := TagsInTodoResponseModel{
		TodoId: tags.TodoId(),
		TagIds: responseTagIds,
	}
	// Response in JSON format
	return ctx.JSON(http.StatusOK, response)
}

// GET Todo項目（複数）の参照
func (h *TodoHandler) GetTodos(ctx echo.Context) error {
	//Request
	var params GetTodosParams
	var trm TodosResponseModel
	var response []TodosResponseModel
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Validation
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.FindTodosUseCaseInputDto{
		Id:     params.Id,
		TodoId: params.TodoId,
		Title:  params.Title,
	}
	// Usecase
	todos, err := todoDi.FindTodos().FindMultiple(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	for _, todo := range todos {
		responseTagIds := make([]string, 0)
		input_dto_for_get_tags := todoApp.GetTagsInTodoUseCaseInputDto{
			TodoId: todo.TodoId(),
		}
		tags, err := todoDi.GetTagsInTodo().GetTagsInTodo(ctx.Request().Context(), input_dto_for_get_tags)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		for _, tagId := range tags.TagIds() {
			responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
		}
		trm = TodosResponseModel{
			Id:          todo.Id(),
			TodoId:      todo.TodoId(),
			Title:       todo.Title(),
			Description: todo.Description(),
			IsDeletable: todo.IsDeletable(),
			IsChecked:   false,
			TagIds:      responseTagIds,
			CreatedAt:   todo.CreatedAt(),
			UpdatedAt:   todo.UpdatedAt(),
		}
		response = append(response, trm)
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}

// Post 新規作成
// 一度に一つしかtodo項目が作成されない想定
func (h *TodoHandler) PostTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostTodosParams
	var responseTagIds []string
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
	// Usecase
	todo, err := todoDi.CreateTodo().Create(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.TagIds {
		tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		tagids_for_use_case = append(tagids_for_use_case, tagid_for_use_case)
	}
	input_dto_for_add_tags := todoApp.AddTagsInTodoUseCaseInputDto{
		TodoId: todo.TodoId(),
		TagIds: tagids_for_use_case,
	}
	tags, err := todoDi.AddTagsInTodo().AddTagsInTodo(ctx.Request().Context(), input_dto_for_add_tags)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		IsChecked:   false,
		TagIds:      responseTagIds,
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}

// PUT 更新
// 一度に一つしかtodo項目が更新されない想定
func (h *TodoHandler) PutTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutTodosParams
	var responseAddedTagIds []string
	var responseDeletedTagIds []string
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	fmt.Printf("params: %v\n", params)
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	post_tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.PostTagIds {
		post_tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		post_tagids_for_use_case = append(post_tagids_for_use_case, post_tagid_for_use_case)
	}
	delete_tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.DeleteTagIds {
		delete_tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		delete_tagids_for_use_case = append(delete_tagids_for_use_case, delete_tagid_for_use_case)
	}
	input_dto := todoApp.UpdateTodoUseCaseInputDto{
		Id:          params.Id,
		TodoId:      params.TodoId,
		Title:       params.Title,
		Description: params.Description,
		IsDeletable: params.IsDeletable,
	}
	input_dto_for_add_tags := todoApp.AddTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
		TagIds: post_tagids_for_use_case,
	}
	input_dto_for_delete_tags := todoApp.DeleteTagsInTodoUseCaseInputDto{
		TodoId: params.TodoId,
		TagIds: delete_tagids_for_use_case,
	}
	// Usecase
	todo, err := todoDi.UpdateTodo().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	addedTags, err := todoDi.AddTagsInTodo().AddTagsInTodo(ctx.Request().Context(), input_dto_for_add_tags)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	deletedTags, err := todoDi.DeleteTagsInTodo().DeleteTagsInTodo(ctx.Request().Context(), input_dto_for_delete_tags)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	for _, tagId := range addedTags.TagIds() {
		responseAddedTagIds = append(responseAddedTagIds, fmt.Sprint(tagId))
	}
	for _, tagId := range deletedTags.TagIds() {
		responseDeletedTagIds = append(responseDeletedTagIds, fmt.Sprint(tagId))
	}
	response := PutTodosResponseModel{
		Id:            todo.Id(),
		TodoId:        todo.TodoId(),
		Title:         todo.Title(),
		Description:   todo.Description(),
		IsDeletable:   todo.IsDeletable(),
		IsChecked:     false,
		AddedTagIds:   responseAddedTagIds,
		DeletedTagIds: responseDeletedTagIds,
		CreatedAt:     todo.CreatedAt(),
		UpdatedAt:     todo.UpdatedAt(),
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}

// DELETE 削除
// 一度に一つしかtodo項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *TodoHandler) DeleteTodo(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTodoParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		fmt.Printf("An error is happening: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := todoApp.DeleteTodoUseCaseInputDto{
		Id:          params.Id,
		TodoId:      params.TodoId,
		IsDeletable: params.IsDeletable,
	}
	// Usecase
	todo, err := todoDi.DeleteTodo().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	// UseCase → Presentation
	response := TodosResponseModel{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		IsChecked:   true,
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) DeleteTodos(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTodosParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		fmt.Printf("An error is happening: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Transform params into dto.
	TodosForDto := make([]todoDomain.TodosForDto, 0)
	for _, todo := range params.Todos {
		TodosForDto = append(TodosForDto, todoDomain.TodosForDto{TodoId: todo.TodoId, IsDeletable: todo.IsDeletable})
	}
	//  Presentation -> UseCase
	input_dto := todoApp.DeleteTodosUseCaseInputDto{
		Id:    params.Id,
		Todos: TodosForDto,
	}
	// Usecase
	todo_result, err := todoDi.DeleteTodos().DeleteTodos(ctx.Request().Context(), input_dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	// Transform the result into response.
	TodosForResponse := make([]struct {
		TodoId      string `json:"todoId"`
		IsDeletable bool   `json:"isDeletable"`
	}, 0)
	// tre(todo result element)
	for _, tre := range todo_result.Todos() {
		tmp := struct {
			TodoId      string `json:"todoId"`
			IsDeletable bool   `json:"isDeletable"`
		}{
			TodoId:      tre.TodoId,
			IsDeletable: tre.IsDeletable,
		}
		TodosForResponse = append(TodosForResponse, tmp)
	}
	// UseCase → Presentation
	response := DeleteTodosResponseModel{
		Id:    todo_result.Id(),
		Todos: TodosForResponse,
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}
