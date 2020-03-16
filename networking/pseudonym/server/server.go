package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/gitcrawler"
)

type Response struct {
	Message string
	Err     error
}

func all(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
	database.Query()
}

func addEntries(w http.ResponseWriter, r *http.Request) {
	response := Response{"", nil}
	httpReturnCode := http.StatusOK

	if r.Method == "POST" {

		randomGitusers, getRandomGitusersErr := gitcrawler.GetRandomUsers()
		if getRandomGitusersErr != nil {
			response.Message = "Failed to get random users..."
			response.Err = getRandomGitusersErr
		} else {
			for _, randomGitusers := range *randomGitusers {
				newPseudonym := &database.Pseudonym{
					Username: randomGitusers.Login,
				}

				newPseudonymDetails := &database.PseudonymDetails{
					HTMLURL:   randomGitusers.HTMLURL,
					Likes:     0,
					Pseudonym: newPseudonym,
				}

				fmt.Println("Inserting user " + newPseudonym.Username)
				userErr := database.InsertUserIntoTable(newPseudonym)
				if userErr != nil {
					response.Message = "Failed to insert user " + newPseudonym.Username
					response.Err = userErr
				} else {
					response.Message = "YAY! Inserted user " + newPseudonym.Username
				}

				fmt.Println("Inserting user detail for user " + newPseudonymDetails.Pseudonym.Username)
				userDetailsErr := database.InsertUserDetailIntoTable(newPseudonymDetails)
				if userDetailsErr != nil {
					response.Message = "Failed to insert user " + newPseudonym.Username
					response.Err = userDetailsErr
				} else {
					response.Message = "YAY! Inserted user " + newPseudonym.Username
				}
			}
		}

	} else {
		response.Message = "Sorry, GET method not accepted here :["
		httpReturnCode = http.StatusBadRequest
	}

	w.WriteHeader(httpReturnCode)

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
