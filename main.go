package main

import (
	"net/http"
	"html/template"


)

type page struct {
	Title string	//имена переменных должны совпадать с тем, что мы написали выше!
	Msg string		//и переменные обязательно должны быть публичными!
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")


	t, _ :=template.ParseFiles("templates/index.html")
	t.Execute(w, &page{Title: "Just page", Msg: "Hello World"})


}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}