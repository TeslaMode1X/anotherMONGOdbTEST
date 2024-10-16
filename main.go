package main

import (
	"fmt"
	"log"
	"mongoAPI/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	fmt.Println("listening at port 4040")
	log.Fatal(http.ListenAndServe(":4040", r))
}
