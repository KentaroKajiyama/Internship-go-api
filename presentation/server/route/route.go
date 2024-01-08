package route

import (
	todoPre "github.com/KentaroKajiyama/internship-go-api/presentation/todo"
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
	h := todoPre.NewTodoHandler()
	todos := routeGroup.Group("/todos")
	todos.POST("/", h.PostTodos)
	todos.PUT("/:todo_id", h.PutTodos)
	todos.DELETE("/:todo_id", h.DeleteTodos)
}
