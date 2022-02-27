package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

type badAuth struct {
	Username string
	Password string
}

func (b *badAuth)ServeHTTP(w http.ResponseWriter,r *http.Request,next http.HandlerFunc)  {
	username :=r.URL.Query().Get("username")
	password :=r.URL.Query().Get("password")
	if username != b.Username || password !=b.Password{
		http.Error(w,"Unauthorized",401)
		return
	}
	ctx := context.WithValue(r.Context(),"username",username)
	r = r.WithContext(ctx)
	next(w,r)
}

func hello(w http.ResponseWriter,r *http.Request)  {
	username :=r.Context().Value("username").(string)
	fmt.Fprintf(w,"Hi %s\n",username)
}


func main() {
	r :=mux.NewRouter()
	r.HandleFunc("/hello",hello).Methods("GET")
	n :=negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "s3cr3ct",
	})
	n.UseHandler(r)
	http.ListenAndServe(":8080",n)

}
