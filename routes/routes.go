package root

import (
	"forum/controller"
	"net/http"
)

func Route() {
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	http.HandleFunc("/", controller.Login)
	http.HandleFunc("/home", controller.Index)
	http.HandleFunc("/home2", controller.Home2)
	http.HandleFunc("/comment", controller.Comment)
	http.HandleFunc("/register", controller.Register)
}
