package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/gitcrawler"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Message string
	Err     error
}

func all(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
}

func addPseudonym(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := Response{"", nil}
	httpReturnCode := http.StatusOK

	if err := verifyMethodType(r.Method, "POST", &response, &w); err == nil {
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
					response.Message = fmt.Sprintf("Failed to insert user %s", newPseudonym.Username)
					response.Err = userErr
				} else {
					response.Message = fmt.Sprintf("YAY! Inserted user %s", newPseudonym.Username)
				}

				fmt.Println("Inserting user detail for user " + newPseudonymDetails.Pseudonym.Username)
				newPseudonymDetails.PseudonymID = newPseudonym.ID
				userDetailsErr := database.InsertUserDetailIntoTable(newPseudonymDetails)
				if userDetailsErr != nil {
					response.Message = fmt.Sprintf("Failed to insert user %s", newPseudonym.Username)
					response.Err = userDetailsErr
				} else {
					response.Message = fmt.Sprintf("YAY! Inserted user %s", newPseudonym.Username)
				}
			}
		}
		w.WriteHeader(httpReturnCode)
	}

	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getPseudonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := Response{"", nil}
	httpReturnCode := http.StatusOK

	if err := verifyMethodType(r.Method, "GET", &response, &w); err == nil {
		w.WriteHeader(httpReturnCode)
		givenIdInt, convertGivenIdErr := strconv.Atoi(ps.ByName("id"))
		if convertGivenIdErr != nil {
			fmt.Println("Error converting the given id to an integer!")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			pseudonym, queryErr := database.Query(givenIdInt)
			if queryErr != nil {
				fmt.Println("Error querying id " + string(givenIdInt))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				response.Message = fmt.Sprintf("%v", pseudonym)
				response.Err = nil
			}

		}

	}
	js, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func verifyMethodType(givenMethod, acceptedMethod string, response *Response, w *http.ResponseWriter) error {
	if strings.Compare(givenMethod, acceptedMethod) == 0 {
		return nil
	} else {
		response.Message = fmt.Sprintf("Sorry, given method %s is not accepted.", givenMethod)
		(*w).WriteHeader(http.StatusBadRequest)
		return errors.New(response.Message)
	}
}

func StartServer() {
	router := httprouter.New()
	router.GET("/all", all)
	router.POST("/addPseudonym", addPseudonym)
	router.GET("/pseudonym/:id", getPseudonym)

	http.ListenAndServe(":8080", router)
}
