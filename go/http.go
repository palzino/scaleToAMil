package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func newRouter() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, rand.Intn(100))
	})
	return r
}
func main() {
	router := newRouter()
	fmt.Println("Starting server on :4040")
	err := http.ListenAndServe(":4040", router)
	if err != nil {
		panic(err)
	}

}
