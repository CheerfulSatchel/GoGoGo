package server

import (
	"fmt"
	"io"
	"net/http"
)

func all(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
}

func StartServer() {
	http.HandleFunc("/all", all)

	http.ListenAndServe(":8080", nil)
}
