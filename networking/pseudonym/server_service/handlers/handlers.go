package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database"
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/gitcrawler"
	"github.com/julienschmidt/httprouter"
)

type Response map[string]interface{}

func AllPseudonyms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
}

func createResponse(message string, err error, data interface{}) *Response {
	return &Response{
		"message": message,
		"err":     err,
		"data":    data,
	}
}

func (response *Response) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	return encoder.Encode(response)
}

func CreatePseudonyms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	verifyMethodTypeErr := verifyMethodType(r.Method, http.MethodPost, &w)

	if verifyMethodTypeErr != nil {
		http.Error(w, verifyMethodTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	randomGitusers, getRandomGitusersErr := gitcrawler.GetRandomUsers()
	if getRandomGitusersErr != nil {
		http.Error(w, getRandomGitusersErr.Error(), http.StatusFailedDependency)
		return
	}

	addedUsers := make([]*database.PseudonymDetails, 2)

	for _, randomGitusers := range *randomGitusers {
		newPseudonym := &database.Pseudonym{
			Username: randomGitusers.Login,
		}

		newPseudonymDetails := &database.PseudonymDetails{
			HTMLURL:   randomGitusers.HTMLURL,
			Likes:     0,
			Pseudonym: newPseudonym,
		}

		fmt.Printf("Inserting user %v\n", newPseudonym.Username)
		userErr := database.InsertUserIntoTable(newPseudonym)
		if userErr != nil {
			fmt.Printf("Failed to insert user %v\n", newPseudonym.Username)
		} else {
			fmt.Printf("YAY! Inserted user %v\n", newPseudonym.Username)
		}

		fmt.Printf("Inserting user detail for user %v\n", newPseudonymDetails.Pseudonym.Username)
		newPseudonymDetails.PseudonymID = newPseudonym.ID
		userDetailsErr := database.InsertUserDetailIntoTable(newPseudonymDetails)
		if userDetailsErr != nil {
			fmt.Printf("Failed to insert user %v\n", newPseudonym.Username)
		} else {
			fmt.Printf("YAY! Inserted user %v\n", newPseudonym.Username)
			addedUsers = append(addedUsers, newPseudonymDetails)
		}
	}

	response := createResponse(
		fmt.Sprintf("Successfully created %v users", len(addedUsers)),
		nil,
		addedUsers)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	toJSONErr := response.ToJSON(w)

	if toJSONErr != nil {
		http.Error(w, "Failure to encode JSON", http.StatusInternalServerError)
		return
	}
}

func GetPseudonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	verifyMethodTypeErr := verifyMethodType(r.Method, http.MethodGet, &w)

	if verifyMethodTypeErr != nil {
		http.Error(w, verifyMethodTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	givenIDInt, convertGivenIDErr := strconv.Atoi(ps.ByName("id"))
	if convertGivenIDErr != nil {
		http.Error(w, convertGivenIDErr.Error(), http.StatusBadRequest)
		return
	}

	pseudonymDetails, queryErr := database.Query(givenIDInt)
	if queryErr != nil {
		http.Error(w, queryErr.Error(), http.StatusBadRequest)
		return
	}

	response := createResponse(
		fmt.Sprintf("Successfully retrieved %v", pseudonymDetails.Pseudonym.Username),
		nil,
		pseudonymDetails)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	toJSONErr := response.ToJSON(w)

	if toJSONErr != nil {
		http.Error(w, "Failure to encode JSON", http.StatusInternalServerError)
		return
	}

}
