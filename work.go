package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<html>
  <form action="/multipartform?lang=cpp&name=dj" method="post" enctype="multipart/form-data">
	  <label>MultipartForm:</label>
      <input type="text" name="lang" />
      <input type="text" name="age" />
      <input type="file" name="uploaded" />
    <button type="submit">提交</button>
  </form>
</html>
`)
}

func multipartFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("Open failed: ", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err == nil {
		fmt.Fprintln(w, string(data))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handle)
	mux.HandleFunc("/multipartform", multipartFormHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
