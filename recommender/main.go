package main

import (
	"./config"
	"./controllers"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	conf := config.New()
	router := routing.New()
	router.Get("/recommendation/<idClient>/<idProduct>", controllers.GetRecommentation)
	panic(fasthttp.ListenAndServe(conf.ServerAddress, router.HandleRequest))
}
