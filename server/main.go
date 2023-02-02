package main

import (
	"fmt"
	"net/http"
	"time"
)

func content(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hey world")
	case "/ninja":
		fmt.Fprint(w, "Naruto")
	case "/samurai":
		fmt.Fprint(w, "jack")
	default:
		fmt.Fprint(w, "Big Fat error!!")
	}
}
func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did *not* timeout")
}
func main() {
	http.HandleFunc("/", content)
	http.HandleFunc("/timeout", timeout)
	http.ListenAndServe(":8081", nil)
	// server := http.Server{
	// 	Addr:         "",
	// 	Handler:      nil,
	// 	ReadTimeout:  1000,
	// 	WriteTimeout: 1000,
	// }
	// server.ListenAndServe()
}
