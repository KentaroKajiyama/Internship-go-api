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
	apiWithId := routeGroup.Group("/api/:id")
	api.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome")
	})
	{
		userRoute(api)
		todoRoute(apiWithId)
		tagRoute(apiWithId)
	}
}

func userRoute(routeGroup *echo.Group) {
	h := userPre.NewUserHandler()
	users := routeGroup.Group("/users")
	users.GET("/:firebase_uid", h.GetUser)
	users.POST("", h.PostUsers)
	users.PUT("/:id", h.PutUsers)
	users.DELETE("/:id", h.DeleteUsers)
}

func todoRoute(routeGroup *echo.Group) {
	h := todoPre.NewTodoHandler()
	todos := routeGroup.Group("/todos")
	todos.GET("/:todo_id", h.GetTodo)
	todos.GET("/:todo_id/tags", h.GetTagsInTodo)
	todos.POST("/:todo_id/tags", h.PostTagsInTodo)
	todos.DELETE("/:todo_id/tags", h.DeleteTagsInTodo)
	todos.GET("", h.GetTodos)
	todos.POST("", h.PostTodos)
	todos.PUT("/:todo_id", h.PutTodos)
	todos.DELETE("/:todo_id", h.DeleteTodo)
	todos.DELETE("", h.DeleteTodos)
}

func tagRoute(routeGroup *echo.Group) {
	h := tagPre.NewTagHandler()
	tags := routeGroup.Group("/tags")
	tags.GET("/:tag_id", h.GetTag)
	tags.GET("", h.GetTagsByTodoId)
	tags.POST("", h.PostTags)
	tags.PUT("/:tag_id", h.PutTag)
	tags.DELETE("/:tag_id", h.DeleteTag)
	tags.DELETE("", h.DeleteTags)
}
