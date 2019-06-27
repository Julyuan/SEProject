package main

import (
	"SecondhandBookTradeWebsite/src/data"
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

func authenticate(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	err := request.ParseForm()
	user, err := data.UserByUid(request.PostFormValue("username"))
	fmt.Print("Uid ",user.Uid, " Upassword ",user.Upassword,"\n")
	if err != nil{
		danger(err, "Cannot find user")
	}
	if user.Upassword == request.PostFormValue("password"){
		session, err := user.CreateSession()
		if err != nil{
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		fmt.Println("session uuid "+session.Uuid)
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	}else{
		http.Redirect(writer,request,"/login",302)
	}
}

func register(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	generateHTML("register_page",writer,nil, "register_page","layouts/footer","layouts/title")
}

func register_verify(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	err := request.ParseForm()
	if err != nil{
		danger(err, "Cannot parse form")
	}

	user := data.User{
		Uid:		request.PostFormValue("username"),
		Upassword:	request.PostFormValue("password"),
	}

	//fmt.Print(user.Uid," ",user.Upassword,"\n")

	if err := user.Create(); err != nil{
		danger(err, "Cannot create user")
		fmt.Fprintf(writer,`{
			"status":"fail"
		}`)
	} else{
		fmt.Fprintf(writer,	`{
			"status":"success"	
		}`)
	}

//	fmt.Println(request.PostFormValue("username"))
//	fmt.Println(request.PostFormValue("password"))
//	fmt.Fprintf(writer,`{
//	"employees": [
//	{ "firstName":"Bill" , "lastName":"Gates" },
//	{ "firstName":"George" , "lastName":"Bush" },
//	{ "firstName":"Thomas" , "lastName":"Carter" }
//]
//}`)
//	fmt.Println(request.URL)
}