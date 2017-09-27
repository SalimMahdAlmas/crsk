package main

import (
	"net/http"
	"log"
	"fmt"
)


func main()	{
	http.HandleFunc("/",	index)
	http.HandleFunc("/login",login)
	err	:=	http.ListenAndServe(":9090",	nil)
	if err	!=	nil	{
		log.Fatal("Error:	",	err)
	}
}
func login(writer http.ResponseWriter, request *http.Request) {

}
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Hello World")
}
