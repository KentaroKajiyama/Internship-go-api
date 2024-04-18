package tag

// validationのctx.Vaildate(params)→ctx.Vaildate(&params)に変更している
import (
	"fmt"
	"net/http"
	"strconv"

	tagApp "github.com/KentaroKajiyama/Internship-go-api/application/tag"
	tagDi "github.com/KentaroKajiyama/Internship-go-api/di/tag"
	"github.com/labstack/echo/v4"
)

type tagHandler struct {
}

func NewTagHandler() *tagHandler {
	return &tagHandler{}
}

// Get tag項目の参照
func (h *tagHandler) GetTag(ctx echo.Context) error {
	//リクエストパラメーター取得
	var params GetTagParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	tagid_for_use_case, err := strconv.ParseUint(params.TagId, 10, 64)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	input_dto := tagApp.FindTagUseCaseInputDto{
		Id:    params.Id,
		TagId: tagid_for_use_case,
	}
	// UseCase処理
	tag, err := tagDi.FindTag().Find(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TagsResponseModel{
		Id:        tag.Id(),
		TagId:     fmt.Sprint(tag.TagId()),
		Name:      tag.Name(),
		IsChecked: false,
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

func (h *tagHandler) GetTagsByTodoId(ctx echo.Context) error {
	// リクエストパラメータ取得
	var params GetTagsParams
	var response []TagsResponseModel
	var trm TagsResponseModel
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// バリデーション
	if err = ctx.Validate(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	input_dto := tagApp.FindTagsByTodoIdUseCaseInputDto{
		Id:     params.Id,
		TodoId: params.TodoId,
		Name:   params.Name,
	}
	// UseCase処理
	tags, err := tagDi.FindTagsByTodoId().FindByTodoId(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	for _, tag := range tags {
		trm = TagsResponseModel{
			Id:        tag.Id(),
			TagId:     fmt.Sprint(tag.TagId()),
			Name:      tag.Name(),
			IsChecked: false,
			CreatedAt: tag.CreatedAt(),
			UpdatedAt: tag.UpdatedAt(),
		}
		response = append(response, trm)
	}
	return ctx.JSON(http.StatusOK, response)
}

// Post 新規作成
// 一度に一つしかtag項目が作成されない想定
func (h *tagHandler) PostTags(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostTagsParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := tagApp.CreateTagUseCaseInputDto{
		Id:   params.Id,
		Name: params.Name,
	}
	// UseCase処理 ここでdbが挿入される
	tag, err := tagDi.CreateTag().Create(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TagsResponseModel{
		Id:        tag.Id(),
		TagId:     fmt.Sprint(tag.TagId()),
		Name:      tag.Name(),
		IsChecked: false,
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// PUT 更新
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtag項目が更新されない想定
func (h *tagHandler) PutTag(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutTagsParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	tagid_for_use_case, err := strconv.ParseUint(params.TagId, 10, 64)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	input_dto := tagApp.UpdateTagUseCaseInputDto{
		Id:    params.Id,
		TagId: tagid_for_use_case,
		Name:  params.Name,
	}
	// UseCase処理
	tag, err := tagDi.UpdateTag().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TagsResponseModel{
		Id:        tag.Id(),
		TagId:     fmt.Sprint(tag.TagId()),
		Name:      tag.Name(),
		IsChecked: false,
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// DELETE 削除
// dtoの部分をどうするか？とりあえず、wireは使わずに直感的に書いてみる
// 一度に一つしかtag項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *tagHandler) DeleteTag(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTagParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	if err = ctx.Validate(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	tagid_for_use_case, err := strconv.ParseUint(params.TagId, 10, 64)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	input_dto := tagApp.DeleteTagUseCaseInputDto{
		Id:    params.Id,
		TagId: tagid_for_use_case,
	}
	// UseCase処理
	tag, err := tagDi.DeleteTag().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := TagsResponseModel{
		Id:        tag.Id(),
		TagId:     fmt.Sprint(tag.TagId()),
		Name:      tag.Name(),
		IsChecked: false,
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

func (h *tagHandler) DeleteTags(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteTagsParams
	var responseTagIds []string
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
	tagids_for_use_case := make([]uint64, 0)
	for _, tagId := range params.TagIds {
		tagid_for_use_case, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		tagids_for_use_case = append(tagids_for_use_case, tagid_for_use_case)
	}
	input_dto := tagApp.DeleteTagsUseCaseInputDto{
		Id:     params.Id,
		TagIds: tagids_for_use_case,
	}
	// Usecase
	tags, err := tagDi.DeleteTags().DeleteTags(ctx.Request().Context(), input_dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	// UseCase → Presentation
	for _, tagId := range tags.TagIds() {
		responseTagIds = append(responseTagIds, fmt.Sprint(tagId))
	}
	response := DeleteTagsResponseModel{
		Id:     tags.Id(),
		TagIds: responseTagIds,
	}
	//Response
	return ctx.JSON(http.StatusOK, response)
}
