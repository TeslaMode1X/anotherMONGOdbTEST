package router

import "net/http"

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to MongoDBApi</h1>"))
}
