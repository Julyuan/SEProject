package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main(){
	p("OnlineStore",version(),"started at", config.Address)

	mux := httprouter.New()
	//files := http.FileServer(http.Dir(config.Static))
	//
	//fmt.Println(http.Dir(config.Static))
	//mux.Handler("get","/static/",http.StripPrefix("/static/",files))
	mux.ServeFiles("/static/*filepath",http.Dir(config.Static))
	mux.GET("/",index)
	mux.GET("/login",login)
	mux.POST("/login/authenticate", authenticate)
	mux.GET("/register",register)
	mux.POST("/",handle_post)
	mux.POST("/register/verify",register_verify)
	//fmt.Println("fkx")
	server := http.Server{
		Addr : "0.0.0.0:8080",
		Handler:mux,
	}

	server.ListenAndServe()
}

func handle_post(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	//fmt.Println(request.)
}