package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database"
)

type Response struct {
	Message string
}

func all(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
	database.Query()
}

func addEntries(w http.ResponseWriter, r *http.Request) {
	response := Response{""}
	if r.Method == "POST" {
		response.Message = "Yay! Added new users to the database!"
		w.WriteHeader(http.StatusOK)
	} else {
		response.Message = "Sorry, GET method not accepted here :["
		w.WriteHeader(http.StatusBadRequest)
	}

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartServer() {
	http.HandleFunc("/all", all)
	http.HandleFunc("/add", addEntries)

	http.ListenAndServe(":8080", nil)
}
