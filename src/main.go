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
	mux.GET("/search", search)
	mux.POST("/authenticate", authenticate)
	mux.GET("/register",register)
	mux.POST("/",handle_post)
	mux.POST("/register/verify",register_verify)
	mux.GET("/detail", detail)
	mux.GET("/myinfo", myinfo)
	mux.GET("/shopinfo",shopinfo)
	mux.GET("/paysuccess",paysuccess)
	mux.GET("/mypaid",mypaid)
	mux.GET("/pay",pay)
	mux.GET("/myorder",myorder)
	mux.GET("/cart",cart)
	mux.GET("/shopmanage",shopmanage)
	mux.GET("/writeorder",writeorder)
	mux.GET("/shopupload",shopupload)
	mux.GET("/locateorder_comment",locateorder_comment)
	mux.GET("/locateorder_finished",locateorder_finished)
	mux.GET("/locateorder_pay",locateorder_pay)
	mux.GET("/locateorder_receive",locateorder_receive)

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