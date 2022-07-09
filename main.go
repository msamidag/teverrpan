package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8080", nil)
}
func anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "MERHABA HEROKU")
}
