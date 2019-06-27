package main

import (
	"SecondhandBookTradeWebsite/src/data"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)
//import "SecondhandBookTradeWebsite/src/data"

type IndexInfo struct {
	Booksnapshots []data.BookSnapshot
	User           data.User
	Test		   int
}

type SearchInfo struct{
	User data.User
	Booksnapshots []data.BookSnapshot
	Searchquery SearchQuery
}
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
	index_info := IndexInfo{}
	sess, err := session(writer, request)
	index_info.Booksnapshots, _ = data.GetRandom8Books()
	//fmt.Println(index_info)
	//fmt.Println()
	if err != nil {
		//fmt.Println(index_info.Test)
		//fmt.Println("没有登录")
		generateHTML("index_page",writer,index_info,"index_page","layouts/footer","layouts/navi",
			"layouts/public.navbar","layouts/header","layouts/index_main","layouts/type_navi","layouts/hotshow")
	} else {
		//fmt.Println("已经登录")
		index_info.User,_ = sess.User()
		generateHTML("index_page",writer,index_info,"index_page","layouts/footer","layouts/navi",
			"layouts/private.navbar","layouts/header","layouts/index_main","layouts/type_navi","layouts/hotshow")
	}

}

type SearchQuery struct {
	Key string
	Booktype string
	Refer string
	Way string
	Dowhat string
	PageCount int
	PageNow int
	PageFrom int
	PageTo int
}

type SearchResult struct {

}

func search(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	fmt.Println("search 网页填充")

	search_info := SearchInfo{}
	sess, err := session(writer, request)

	url_queries := request.URL.Query()
	search_query := SearchQuery{}
	search_query.Dowhat	= url_queries["dowhat"][0]

	if(search_query.Dowhat=="searchByRand"){

	}else{
		search_query.Key = url_queries["key"][0]
		search_query.Dowhat = "searchByKey"
		//search_query.Refer = url_queries["refer"][0]
		//search_query.Way = url_queries["way"][0]
		//search_query.Booktype = url_queries["type"][0]
		//search_query.PageCount,_ = strconv.Atoi(url_queries["pageCount"][0])
		//search_query.PageNow,_ = strconv.Atoi(url_queries["pageNow"][0])
		//search_query.PageFrom,_ = strconv.Atoi(url_queries["pageFrom"][0])
		//search_query.PageTo,_ = strconv.Atoi(url_queries["pageTo"][0])
		if(search_query.Dowhat=="searchByKey"){
			search_query.PageCount,_ = data.GetBookSnapshotsPageCountByKey(search_query.Key);
		}else if(search_query.Dowhat == "searchByType"){
			search_query.PageCount,_ = data.GetBookSnapshotsPageCountByType(search_query.Booktype);
		}

		//if(search_query.Way=="byPageLast"){
		//	if(search_query.PageNow == search_query.PageFrom) {
		//		search_query.PageFrom -= 1
		//	} else{
		//		search_query.PageFrom = search_query.PageFrom;
		//	}
		//	search_query.PageNow--;
		//}else if(search_query.Way=="byPageNext"){
		//	if(search_query.PageNow==search_query.PageTo){
		//		search_query.PageFrom+=1;
		//	}else{
		//
		//	}
		//}else if(search_query.Way=="byPageJump"){
		//	if(search_query.PageFrom+search_query.PageNow-search_query.PageTo<1){
		//		search_query.PageFrom = 1
		//	}else {
		//		search_query.PageFrom =  search_query.PageFrom+search_query.PageNow-search_query.PageTo
		//	}
		//}

		//if(search_query.PageCount<=data.PageRange){
		//	search_query.PageTo = search_query.PageCount
		//}else {
		//	search_query.PageTo = search_query.PageFrom+data.PageRange-1
		//}
		if search_query.Dowhat=="searchByKey" {
			search_info.Booksnapshots, _ = data.GetBookSnapshotsByKey(search_query.Key, search_query.PageNow, search_query.Refer)
		}else if search_query.Dowhat=="searchByType"{
			search_info.Booksnapshots,_ = data.GetBookSnapshotsByType(search_query.Booktype,search_query.PageNow,search_query.Refer)
		}
		fmt.Println(search_info.Booksnapshots)
	}
	//search_info.Searchquery = search_query
	//fmt.Println(url_queries)
	if err != nil{
		//fmt.Println("没有登录")
		//search_info.User,_ = sess.User()
		generateHTML("search_page",writer,search_info,"search_page","layouts/footer","layouts/navi",
			"layouts/public.navbar","layouts/header","layouts/search_main","layouts/orderbar","layouts/searchshow",
			"layouts/search_pagebar")
	}else{
		//fmt.Println("已经登录")
		search_info.User,_ = sess.User()
		generateHTML("search_page",writer,search_info,"search_page","layouts/footer","layouts/navi",
			"layouts/private.navbar","layouts/header","layouts/search_main","layouts/orderbar","layouts/searchshow",
			"layouts/search_pagebar")
	}
}

type DetailQuery struct {
	Bid string
	Dowhat string
}

type detail_info struct{
	BookDetail data.BookDetail
	BookCommentPageCount int
	BookTradeRecordPageCount int

}


func detail(writer http.ResponseWriter, request *http.Request, p httprouter.Params){
	detail_info := detail_info{}
	//sess, err := session(writer, request)

	url_queries := request.URL.Query()
	detail_query := DetailQuery{}
	//detail_query.Dowhat	= url_queries["dowhat"][0]
	detail_query.Bid = url_queries["Bid"][0]

	//if(detail_query.Dowhat=="findDetail"){
	detail_info.BookDetail,_ = data.GetBookDetailById(detail_query.Bid)
	detail_info.BookCommentPageCount = data.GetPageCount(detail_info.BookDetail.Bcommentnum, data.BookCommentPageSize)
	detail_info.BookTradeRecordPageCount =data.GetPageCount(detail_info.BookDetail.Bsalednum,data.BookTradeRecordPageSize)
	//}else{
	//
	//}
	fmt.Println("网页填充")
	generateHTML("detail_page",writer,detail_info,"detail_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/detail_main","layouts/keydetail","layouts/moredetail")

}

func myinfo(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("myinfo 网页填充")

	//user := data.User{}

	generateHTML("myinfo_page",writer,nil,"myinfo_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/myinfo_main")
}

func shopinfo(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("shopinfo 网页填充")

	//user := data.User{}

	generateHTML("shopinfo_page",writer,nil,"shopinfo_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/shopinfo_main")
}

func paysuccess(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("paysuccess 网页填充")

	//user := data.User{}

	generateHTML("paysuccess_page",writer,nil,"paysuccess_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/paysuccess_main")
}


func mypaid(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("mypaid 网页填充")

	//user := data.User{}

	generateHTML("mypaid_page",writer,nil,"mypaid_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/mypaid_main")
}


func pay(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("pay 网页填充")

	//user := data.User{}

	generateHTML("pay_page",writer,nil,"pay_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/pay_main")
}


func myorder(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("网页填充")

	//user := data.User{}

	generateHTML("myorder_page",writer,nil,"myorder_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/myorder_main")
}


func cart(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("cart 网页填充")

	//user := data.User{}

	generateHTML("cart_page",writer,nil,"cart_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/cart_main")
}


func shopmanage(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("shopmanage 网页填充")

	//user := data.User{}

	generateHTML("shopmanage_page",writer,nil,"shopmanage_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/shopmanage_main")
}


type WriteOrderInfo struct {
	Woderform data.OrderForm
}

func writeorder(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("writeorder 网页填充")

	//user := data.User{}
	write_order_info := WriteOrderInfo{}
	url_queries := request.URL.Query()
	price,_:=strconv.ParseFloat(url_queries["Bprice"][0], 32)
	sid,_:=strconv.Atoi(url_queries["Sid"][0])
	oriprice,_:=strconv.ParseFloat(url_queries["Boriprice"][0], 32)
	num, _ := strconv.Atoi(url_queries["Tboughtnum"][0])
	bookInCart := data.BookInCart{
		Bid:url_queries["Bid"][0],
		Bimage:url_queries["Bid"][0],
		Bprice:price,
		Bname:url_queries["Bname"][0],
		Sid:sid,
		Sicon:url_queries["Sicon"][0],
		Sname:url_queries["Sname"][0],
		Boriprice:oriprice,
		Tboughtnum:num,
		}
	fmt.Println(bookInCart)
	write_order_info.Woderform.BookInCartList=append(write_order_info.Woderform.BookInCartList, bookInCart)

	generateHTML("writeorder_page",writer,write_order_info,"writeorder_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/writeorder_main","layouts/step_bar","layouts/writeorder_address","layouts/writeorder_info")
}


func shopupload(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("shopupload 网页填充")

	//user := data.User{}

	generateHTML("shopupload_page",writer,nil,"shopupload_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/shopupload_main")
}


func locateorder_comment(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("locateorder_comment 网页填充")

	//user := data.User{}

	generateHTML("locateorder_comment_page",writer,nil,"locateorder_comment_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/locateorder_comment_main")
}


func locateorder_finished(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("locateorder_finished 网页填充")

	//user := data.User{}

	generateHTML("locateorder_finished_page",writer,nil,"locateorder_finished_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/locateorder_finished_main")
}


func locateorder_pay(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("locateorder_pay 网页填充")

	//user := data.User{}

	generateHTML("locateorder_pay_page",writer,nil,"locateorder_pay_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/locateorder_pay_main")
}


func locateorder_receive(writer http.ResponseWriter, request *http.Request, p httprouter.Params)  {
	fmt.Println("locateorder_receive 网页填充")

	//user := data.User{}

	generateHTML("locateorder_receive_page",writer,nil,"locateorder_receive_page","layouts/footer","layouts/navi",
		"layouts/public.navbar","layouts/header","layouts/locateorder_receive_main")
}
