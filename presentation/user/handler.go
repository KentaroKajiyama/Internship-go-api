package user

import (
	validator "github.com/KentaroKajiyama/internship-go-api/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/KentaroKajiyama/internship-go-api/di"
	"net/http"
	userUseCase "github.com/KentaroKajiyama/internship-go-api/application/user"
)

type userHandler struct {
	registUserUseCase *userUseCase.RegistUserUseCase
	updateUserUseCase *userUseCase.UpdateUserUseCase
	deleteUserUseCase *userUseCase.DeleteUserUseCase
}

func NewUserHandler(
	registUserUseCase *userUseCase.RegistUserUseCase,
	updateUserUseCase *userUseCase.UpdateUserUseCase,
	deleteUserUseCase *userUseCase.DeleteUserUseCase,
) userHandler {
	return userHandler{
		registUserUseCase: registUserUseCase,
		updateUserUseCase: updateUserUseCase,
		deleteUserUseCase: deleteUserUseCase,
	}
}

func (uh *userHandler) 