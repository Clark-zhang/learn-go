package router

import(
    "github.com/julienschmidt/httprouter"
    "github.com/Clark-zhang/gml_server/controller"
)

func New() *httprouter.Router{
    router := httprouter.New()
    router.GET("/bookId", controller.GetBookId)
    router.POST("/getBook", controller.GetBook)

    return router
}