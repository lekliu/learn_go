package main

import (
	"fmt"
	"net/http"
	"os"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello world"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/posts/Go2/", f2)
	http.ListenAndServe("0.0.0.0:9091", nil)
}
