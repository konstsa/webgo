package main

import (
	"html/template"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"net/http"
)

type page struct {
	Title string //имена переменных должны совпадать с тем, что мы написали выше!
	Msg   string //и переменные обязательно должны быть публичными!
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	title := r.URL.Path[len("/"):]

	if title != "exec/" {
		t, okt := template.ParseFiles("index.html")
		if okt != nil {
			panic(okt)
		}
		t.Execute(w, &page{Title: "Convert Image"})
	} else {
		imgfile, fhead, _ := r.FormFile("imgfile")

		img, ext, ok := image.Decode(imgfile)
		if ok != nil {
			t, okt := template.ParseFiles("index.html")
			if okt != nil {
				panic(okt)
			}
			t.Execute(w, &page{Title: "Convert Image", Msg: ok.Error()})
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
