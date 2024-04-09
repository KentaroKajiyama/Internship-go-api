//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/KentaroKajiyama/Internship-go-api/application/todo"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/google/wire"
)

var provideSetFind = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// Repository
	repository.NewTodoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewFindTodoUseCase,
)

func FindTodo() *todo.FindTodoUseCase {
	wire.Build(
		provideSetFind,
	)
	return nil
}

var provideSetFindMultiple = wire.NewSet(
	//driver
	infrastructure.NewGormPostgres,
	//Repository
	repository.NewTodoRepository,
	//usecase
	todo.NewFindTodosUseCase,
)

func FindTodos() *todo.FindTodosUseCase {
	wire.Build(
		provideSetFindMultiple,
	)
	return nil
}

var provideSetCreate = wire.NewSet(
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
	todo.NewCreateTodoUseCase,
)

func CreateTodo() *todo.CreateTodoUseCase {
	wire.Build(
		provideSetCreate,
	)
	return nil
}

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
	return nil
}

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
	return nil
}
