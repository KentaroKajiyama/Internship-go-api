//go:build wireinject
// +build wireinject

package user

import (
	"github.com/KentaroKajiyama/Internship-go-api/application/user"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/google/wire"
)

var provideSetFind = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// Repository
	repository.NewUserRepository,

	// queryService

	// domainService

	// useCase
	user.NewFindUserUseCase,
)

func FindUser() *user.FindUserUseCase {
	wire.Build(
		provideSetFind,
	)
	return nil
}

var provideSetSignUp = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// // client
	// auth.NewAuthMockClient,
	// // Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	// //auth.NewAuthMock2Client,

	// Repository
	repository.NewUserRepository,

	// queryService

	// domainService

	// useCase
	user.NewSignUpUserUseCase,
)

func SignUpUser() *user.SignUpUserUseCase {
	wire.Build(
		provideSetSignUp,
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
	repository.NewUserRepository,

	// queryService

	// domainService

	// useCase
	user.NewDeleteUserUseCase,
)

func DeleteUser() *user.DeleteUserUseCase {
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
	repository.NewUserRepository,

	// queryService

	// domainService

	// useCase
	user.NewUpdateUserUseCase,
)

func UpdateUser() *user.UpdateUserUseCase {
	wire.Build(
		provideSetUpdate,
	)
	return nil
}
