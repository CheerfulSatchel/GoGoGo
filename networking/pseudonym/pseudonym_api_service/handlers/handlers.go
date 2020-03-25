package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

// func GetPseudonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	verifyMethodTypeErr := verifyMethodType(r.Method, http.MethodGet, &w)

// 	if verifyMethodTypeErr != nil {
// 		http.Error(w, verifyMethodTypeErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	givenIDInt, convertGivenIDErr := strconv.Atoi(ps.ByName("id"))
// 	if convertGivenIDErr != nil {
// 		http.Error(w, convertGivenIDErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	pseudonymDetails, queryErr := database.Query(givenIDInt)
// 	if queryErr != nil {
// 		http.Error(w, queryErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	response := createResponse(
// 		fmt.Sprintf("Successfully retrieved %v", pseudonymDetails.Pseudonym.Username),
// 		nil,
// 		pseudonymDetails)

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	toJSONErr := response.ToJSON(w)

// 	if toJSONErr != nil {
// 		http.Error(w, "Failure to encode JSON", http.StatusInternalServerError)
// 		return
// 	}

// }
