//go:build wireinject
// +build wireinject

package di

// import (
// 	"github.com/KentaroKajiyama/Internship-go-api/application/sample"
// 	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
// 	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/auth"
// 	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/database"
// 	"github.com/google/wire"
// )

// var providerSet = wire.NewSet(
// 	// driver
// 	infrastructure.NewGormPostgres,

// 	// client
// 	auth.NewAuthMockClient,
// 	// Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
// 	//auth.NewAuthMock2Client,

// 	// Repository
// 	database.NewSampleRepository,

// 	// queryService

// 	// domainService

// 	// useCase
// 	sample.NewSamplePingPong,
// )

// func SamplePingPong() *sample.PingPong {
// 	wire.Build(
// 		providerSet,
// 	)
// 	return nil
// }
