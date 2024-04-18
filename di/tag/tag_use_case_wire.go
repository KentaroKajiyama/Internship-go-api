//go:build wireinject
// +build wireinject

package tag

import (
	"github.com/KentaroKajiyama/Internship-go-api/application/tag"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/google/wire"
)

var provideSetFind = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// Repository
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewFindTagUseCase,
)

func FindTag() *tag.FindTagUseCase {
	wire.Build(
		provideSetFind,
	)
	return nil
}

var provideSetFindByTodoId = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// Repository
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewFindTagsByTodoIdUseCase,
)

func FindTagsByTodoId() *tag.FindTagsByTodoIdUseCase {
	wire.Build(
		provideSetFindByTodoId,
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
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewCreateTagUseCase,
)

func CreateTag() *tag.CreateTagUseCase {
	wire.Build(
		provideSetCreate,
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
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewUpdateTagUseCase,
)

func UpdateTag() *tag.UpdateTagUseCase {
	wire.Build(
		provideSetUpdate,
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
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewDeleteTagUseCase,
)

func DeleteTag() *tag.DeleteTagUseCase {
	wire.Build(
		provideSetDelete,
	)
	return nil
}

var provideSetDeleteMultiple = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// // client
	// auth.NewAuthMockClient,
	// // Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	// //auth.NewAuthMock2Client,

	// Repository
	repository.NewTagRepository,

	// queryService

	// domainService

	// useCase
	tag.NewDeleteTagsUseCase,
)

func DeleteTags() *tag.DeleteTagsUseCase {
	wire.Build(
		provideSetDeleteMultiple,
	)
	return nil
}
