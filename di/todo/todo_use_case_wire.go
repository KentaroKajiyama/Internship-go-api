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
	repository.NewToDoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewFindToDoUseCase,
)

func FindToDo() *todo.FindToDoUseCase {
	wire.Build(
		provideSetFind,
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
	repository.NewToDoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewCreateToDoUseCase,
)

func CreateToDo() *todo.CreateToDoUseCase {
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
	repository.NewToDoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewDeleteToDoUseCase,
)

func DeleteToDo() *todo.DeleteToDoUseCase {
	wire.Build(
		provideSetDelete,
	)
	return &todo.DeleteToDoUseCase{}
}

var provideSetUpdate = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// // client
	// auth.NewAuthMockClient,
	// // Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	// //auth.NewAuthMock2Client,

	// Repository
	repository.NewToDoRepository,

	// queryService

	// domainService

	// useCase
	todo.NewUpdateToDoUseCase,
)

func UpdateToDo() *todo.UpdateToDoUseCase {
	wire.Build(
		provideSetUpdate,
	)
	return &todo.UpdateToDoUseCase{}
}
