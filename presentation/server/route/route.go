package route

import (
	todoPre "github.com/KentaroKajiyama/Internship-go-api/presentation/todo"
	"github.com/labstack/echo/v4"
)

func InitRoute(routeGroup *echo.Group) {
	//初めにエラーハンドリングを挟む必要があるかも
	users := routeGroup.Group("/users")
	//最初の認証、ログイン
	//users.GET("/", )
	{
		todoRoute(users)
	}
}

// func userRoute(routeGroup.Group *echo.Group) {

// }

func todoRoute(routeGroup *echo.Group) {
	h := todoPre.NewToDoHandler()
	todos := routeGroup.Group("/todos")
	todos.POST("/", h.PostToDos)
	todos.PUT("/:todo_id", h.PutToDos)
	todos.DELETE("/:todo_id", h.DeleteToDos)
}
