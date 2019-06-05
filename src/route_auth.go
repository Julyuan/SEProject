package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func login(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	//threads, err := data.Threads()
	//if err != nil {
	//	error_message(writer, request, "Cannot get threads")
	//} else {
	//	_, err := session(writer, request)
	//	if err != nil {
	//		generateHTML(writer, threads, "layout", "public.navbar", "index")
	//	} else {
	//		generateHTML(writer, threads, "layout", "private.navbar", "index")
	//	}
	//}
	//generateHTML("index_page",writer,nil,"index_page","layouts/footer","layouts/navi",
	//	"layouts/public.navbar","layouts/header","layouts/index_main","layouts/type_navi")
	generateHTML("login_page",writer,nil,"login_page","layouts/footer","layouts/title");

}


func register(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	generateHTML("register_page",writer,nil, "register_page","layouts/footer","layouts/title")
}

func register_verify(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	fmt.Println(request.PostFormValue("username"))
	fmt.Println(request.PostFormValue("password"))
	fmt.Fprintf(writer,`{
	"employees": [
	{ "firstName":"Bill" , "lastName":"Gates" },
	{ "firstName":"George" , "lastName":"Bush" },
	{ "firstName":"Thomas" , "lastName":"Carter" }
]
}`)
	fmt.Println(request.URL)
}