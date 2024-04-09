package route

import (
	"net/http"

	tagPre "github.com/KentaroKajiyama/Internship-go-api/presentation/tag"
	todoPre "github.com/KentaroKajiyama/Internship-go-api/presentation/todo"
	userPre "github.com/KentaroKajiyama/Internship-go-api/presentation/user"
	"github.com/labstack/echo/v4"
)

func InitRoute(routeGroup *echo.Group) {
	//初めにエラーハンドリングを挟む必要があるかも
	//user_idをもらってくる必要がある。
	api := routeGroup.Group("/api")
	api.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome")
	})
	{
		userRoute(api)
		todoRoute(api)
		tagRoute(api)
	}
}

func userRoute(routeGroup *echo.Group) {
	h := userPre.NewUserHandler()
	users := routeGroup.Group("/users")
	users.GET("/:id", h.GetUsers)
	users.POST("", h.PostUsers)
	users.PUT("/:id", h.PutUsers)
	users.DELETE("/:id", h.DeleteUsers)
}

func todoRoute(routeGroup *echo.Group) {
	h := todoPre.NewTodoHandler()
	todos := routeGroup.Group("/todos")
	todos.GET("/:todo_id", h.GetTodo)
	todos.GET("", h.GetTodos)
	todos.POST("", h.PostTodos)
	todos.PUT("/:todo_id", h.PutTodos)
	todos.DELETE("/:todo_id", h.DeleteTodos)
}

func tagRoute(routeGroup *echo.Group) {
	h := tagPre.NewTagHandler()
	tags := routeGroup.Group("/tags")
	tags.GET("/:tag_id", h.GetTag)
	tags.GET("", h.GetTags)
	tags.POST("", h.PostTags)
	tags.PUT("/:todo_id", h.PutTags)
	tags.DELETE("/:todo_id", h.DeleteTags)
}
