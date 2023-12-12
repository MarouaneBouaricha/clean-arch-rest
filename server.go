package main

import (
	"clean-arch/controller"
	router "clean-arch/http"
	"clean-arch/repository"
	"clean-arch/service"
	"fmt"
	"net/http"
)

var (
	postReposiroty repository.PostRepository = repository.NewPostgresRepository()
	postService    service.PostService       = service.NewPostService(postReposiroty)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
