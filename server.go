package main

import (
	"cleango/controller"
	router "cleango/http"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

var(
	postController controller.PostController = controller.NewPostController()
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "server up")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVER(port)
}
