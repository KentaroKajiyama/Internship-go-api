//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/KentaroKajiyama/internship-go-gpi/application/todo"
	"github.com/KentaroKajiyama/internship-go-gpi/infrastructure"
	"github.com/KentaroKajiyama/internship-go-gpi/infrastructure/repository"
	"github.com/google/wire"
	"kiravia.com/internship-go-api/application/todo"
)

var provideSetDelete = wire.NewSet(
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
	todo.NewDeleteTodoUseCase,
)

func DeleteTodo() *todo.DeleteTodoUseCase {
	wire.Build(
		provideSetDelete,
	)
	return &todo.DeleteTodoUseCase{}
}
