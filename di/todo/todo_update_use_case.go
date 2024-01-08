//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/KentaroKajiyama/Internship-go-api/application/todo"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/google/wire"
)

var provideSetUpdate = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// // client
	// auth.NewAuthMockClient,
	// // Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	// //auth.NewAuthMock2Client,

	// Repository
	repository.NewTodoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewUpdateTodoUseCase,
)

func UpdateTodo() *todo.UpdateTodoUseCase {
	wire.Build(
		provideSetUpdate,
	)
	return &todo.UpdateTodoUseCase{}
}
