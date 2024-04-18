package user

import (
	"fmt"
	"net/http"

	userApp "github.com/KentaroKajiyama/Internship-go-api/application/user"
	userDi "github.com/KentaroKajiyama/Internship-go-api/di/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUser(ctx echo.Context) error {
	//リクエストパラメーター取得（リクエストのボディに対するエラーハンドリング→データ型や形式等が合っているか？）
	var params GetUserParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Errorもろたで")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか→データの内容が特定のバリデーションルールに違反していないか？文字数や書き方など）
	// I don't implement the case of ":id" param, only the case of ":firebase_uid".
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := userApp.FindUserByUidUseCaseInputDto{
		FirebaseUid: params.FirebaseUid,
	}
	// UseCase処理
	user, err := userDi.FindUserByUid().FindByUid(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := UsersResponseModel{
		Id:          user.Id(),
		FirebaseUid: user.FirebaseUid(),
		Name:        user.Name(),
		Email:       user.Email(),
		CreatedAt:   user.CreatedAt(),
		UpdatedAt:   user.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// Post 新規作成
// 一度に一つしかUser項目が作成されない想定
func (h *UserHandler) PostUsers(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PostUsersParams
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := userApp.SignUpUserUseCaseInputDto{
		FirebaseUid: params.FirebaseUid,
		Name:        params.Name,
		Email:       params.Email,
	}
	// UseCase処理 ここでdbが挿入される
	user, err := userDi.SignUpUser().SignUp(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := UsersResponseModel{
		Id:          user.Id(),
		FirebaseUid: user.FirebaseUid(),
		Name:        user.Name(),
		Email:       user.Email(),
		CreatedAt:   user.CreatedAt(),
		UpdatedAt:   user.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// PUT 更新
// 一度に一つしかtodo項目が更新されない想定
func (h *UserHandler) PutUsers(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params PutUsersParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := userApp.UpdateUserUseCaseInputDto{
		Id:    params.Id,
		Name:  params.Name,
		Email: params.Email,
	}
	// UseCase処理
	user, err := userDi.UpdateUser().Update(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := UsersResponseModel{
		Id:          user.Id(),
		FirebaseUid: user.FirebaseUid(),
		Name:        user.Name(),
		Email:       user.Email(),
		CreatedAt:   user.CreatedAt(),
		UpdatedAt:   user.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}

// DELETE 削除
// 一度に一つしかtodo項目が削除されない想定？流石に削除は複数個まとめたい。
func (h *UserHandler) DeleteUsers(ctx echo.Context) error {
	// リクエストパラメーター取得
	var params DeleteUsersParams
	err := ctx.Bind(&params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	//バリデーション（上のerrorハンドリングとはどう違うのか）
	if err = ctx.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//  Presentation -> UseCase
	input_dto := userApp.DeleteUserUseCaseInputDto{
		Id:          params.Id,
		FirebaseUid: params.FirebaseUid,
		Name:        params.Name,
		Email:       params.Email,
	}
	// UseCase処理
	user, err := userDi.DeleteUser().Delete(ctx.Request().Context(), input_dto)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// UseCase → Presentation
	response := UsersResponseModel{
		Id:          user.Id(),
		FirebaseUid: user.FirebaseUid(),
		Name:        user.Name(),
		Email:       user.Email(),
		CreatedAt:   user.CreatedAt(),
		UpdatedAt:   user.UpdatedAt(),
	}
	//レスポンス。JSON形式でいいのか？
	return ctx.JSON(http.StatusOK, response)
}
