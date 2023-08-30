package main

import (
	"fmt"
	handler "forum/controller"
	"net/http"

	"github.com/gofrs/uuid"
)

var u1 = uuid.Must(uuid.NewV4())

func main() {
	fmt.Println(u1)

	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))

	http.HandleFunc("/", handler.Index)

	fmt.Println("Listening in the port 3000...")
	fmt.Println("http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
