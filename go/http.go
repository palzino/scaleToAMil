package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func newRouter() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, strconv.Itoa(rand.Intn(100)))
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
