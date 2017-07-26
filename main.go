package main

import (
	"fmt"
	"html/template"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"net/http"
)

type page struct {
	Title string
	Msg   string
}

func indexpage(w http.ResponseWriter, p map[string]string) {
	w.Header().Set("Content-type", "text/html")
	t, okt := template.ParseFiles("index.html")
	if okt != nil {
		fmt.Println(okt.Error())
		panic(okt)
		return
	}
	t.Execute(w, &page{Title: p["title"], Msg: p["msg"]})
}

func index(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	maps := make(map[string]string)

	if title != "exec/" {
		maps["title"] = "Convert Image"
		indexpage(w, maps)
	} else {
		imgfile, fhead, ok := r.FormFile("imgfile")
		if ok != nil {
			maps["title"] = title
			maps["msg"] = ok.Error()
			indexpage(w, maps)
			return
		}

		img, ext, ok := image.Decode(imgfile)
		if ok != nil {
			maps["title"] = title
			maps["msg"] = ok.Error()
			indexpage(w, maps)
			return
		}
		w.Header().Set("Content-type", "image/jpeg")
		w.Header().Set("Content-Disposition", "filename=\""+fhead.Filename+"."+ext+"\"")
		jpeg.Encode(w, img, &jpeg.Options{0})
	}

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}
