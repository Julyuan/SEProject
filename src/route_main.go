package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
//import "SecondhandBookTradeWebsite/src/data"

func err(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	//vals := request.URL.Query()
	//_, err := session(writer, request)
	//if err != nil {
	//	generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	//} else {
	//	generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	//}
}

func index(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
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
	generateHTML("index_page",writer,nil,"index_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/index_main","layouts/type_navi")
}

