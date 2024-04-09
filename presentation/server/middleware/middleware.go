package middleware

import (
	"context"
	"fmt"
	"net/http"

	"firebase.google.com/go/v4/auth"
	userApp "github.com/KentaroKajiyama/Internship-go-api/application/user"
	userDi "github.com/KentaroKajiyama/Internship-go-api/di/user"
	"github.com/labstack/echo/v4"
)

// FirebaseAuthMiddleware はFirebase Authを使用してリクエストを認証するミドルウェアです。
func FirebaseAuthMiddleware(client *auth.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// 認証トークンをリクエストヘッダーから取得
			idToken := ctx.Request().Header.Get("Authorization")
			if idToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
			}
			// Firebaseでトークンを検証
			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired ID token")
			}
			// Email情報を付与？
			if email, ok := token.Claims["email"].(string); ok && email != "" {
				input_dto := userApp.SignUpUserUseCaseInputDto{
					Name:  email,
					Email: email,
				}
				// User情報を生成してもらう Presentation → Usecase
				user, err := userDi.SignUpUser().SignUp(ctx.Request().Context(), input_dto)
				if err != nil {
					return fmt.Errorf("%w", err)
				}
				ctx.Set("id", user.Id)
				ctx.Set("email", user.Email)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Email is required but not present in the token.")
			}
			return next(ctx)
		}
	}
}
