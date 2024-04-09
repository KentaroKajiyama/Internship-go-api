package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/KentaroKajiyama/Internship-go-api/config"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/validator"
	MyMiddleware "github.com/KentaroKajiyama/Internship-go-api/presentation/server/middleware"
	"github.com/KentaroKajiyama/Internship-go-api/presentation/server/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

func main() {
	// // ToDo: DockerComposeを利用してdatabaseを作成することができたら、以下のコメントアウトを外す
	db := infrastructure.NewGormPostgres()
	defer func() {
		d, _ := db.DB()
		d.Close()
	}()
	//Firebase SDKの初期化
	ctx := context.Background()
	opt := option.WithCredentialsJSON(config.Conf.Google.CredentialsJson)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Firebase Auth インスタンス取得
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// サーバーエンジンの生成
	engine := echo.New()
	engine.Debug = true
	engine.Validator = validator.NewValidator()

	//ここで最後に認証をかます
	engine.Pre(middleware.RemoveTrailingSlash())
	engine.Use(middleware.Recover())
	// baseRoute := engine.Group("")
	// health := engine.Group("/health")
	engine.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OKですよう")
	})

	v1 := engine.Group("/v1")
	v1.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome")
	})
	//認証を入れたらそこからuserIDを取ってくる。
	engine.Use(MyMiddleware.FirebaseAuthMiddleware(client))
	route.InitRoute(v1)

	go func() {
		//ここでのhttp.ErrServerClosedとengine.Shutdown(ctx)とのつながりがわからない。
		if err := engine.Start(fmt.Sprintf(":%s", config.Conf.GetPort())); err != nil && !errors.Is(err, http.ErrServerClosed) {
			engine.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := engine.Shutdown(ctx); err != nil {
		engine.Logger.Fatal(err)
	}
	println("stop server method")
}
