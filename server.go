package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error Parsing Form: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")

	file, handler, err := r.FormFile("myFile")
	fmt.Println(file)
	fmt.Println(handler)
	fmt.Println(err)

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	file_name, err := os.Create(handler.Filename)
	defer file_name.Close()

	n2, err := file_name.Write(fileBytes)

	fmt.Printf("wrote %d bytes\n", n2)

}

func main() {

	fileServer := http.FileServer((http.Dir("./static")))
	http.Handle("/", fileServer)
	http.HandleFunc("/upload", formHandler)

	fmt.Printf("Starting Server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
