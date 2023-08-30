package main

import (
	"fmt"
	route "forum/routes"
	"net/http"

	"github.com/gofrs/uuid"
)

var u1 = uuid.Must(uuid.NewV4())

func main() {
	fmt.Println(u1)
	route.Route()
	fmt.Println("Listening in the port 3000...")
	fmt.Println("http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
