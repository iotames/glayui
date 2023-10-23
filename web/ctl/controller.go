package ctl

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for k, v := range query {
		fmt.Println("----url--query---", k, v)
		fmt.Println("--------------------", r.URL.Host, r.URL.Path)
	}
	w.Write([]byte("This is Home Page"))
}
