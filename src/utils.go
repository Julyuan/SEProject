package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"SecondhandBookTradeWebsite/src/data"
	"errors"
)

type Configuration struct {
	Address string
	ReadTimeout int64
	WriteTimeout int64
	Static string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}){
	fmt.Println(a)
}

func init(){
	loadConfig()
	file, err :=  os.OpenFile("onlinestore.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig(){
	file, err := os.Open("src/config.json")
	if err != nil{
		log.Fatalln("Cannot open config file", err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil{
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func error_message(writer http.ResponseWriter, request *http.Request, msg string){
	utl := []string{"/err?msg=",msg}
	http.Redirect(writer, request, strings.Join(utl,""), 302)
}


func session(writer http.ResponseWriter, request *http.Request)(sess data.Session, err error){
	//cookies := request.Cookies()
	//for _,c := range cookies{
	//	fmt.Println(c.Name)
	//}
	cookie, err := request.Cookie("_cookie")
	if err == nil{
		fmt.Println("err is nil")
		//fmt.Println(cookie.Name+" "+cookie.Value)
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok{
			err = errors.New("Invalid session")
		}
	}else{
		fmt.Println(err)
		//fmt.Println("Cookie初始化出错")
	}
	return
}


func parseTemplateFiles(filenames ...string)(t *template.Template){
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(entry string, writer http.ResponseWriter, data interface{}, filenames ...string){
	var files []string
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("WebContent/html/%s.html", file))

	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer,entry,data)
}

func info(args ...interface{}){
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}){
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}){
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
// version
func version() string {
	return "0.1"
}

