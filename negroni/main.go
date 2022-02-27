package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	r :=mux.NewRouter()
	n :=negroni.Classic()
	n.UseHandler(r)
	http.ListenAndServe(":8080",n)

}
