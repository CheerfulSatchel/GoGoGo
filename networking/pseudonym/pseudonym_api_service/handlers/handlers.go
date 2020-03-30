package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/pseudonym_api_service/models"
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

func CreatePseudonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	verifyMethodTypeErr := verifyMethodType(r.Method, http.MethodPut, &w)

	if verifyMethodTypeErr != nil {
		http.Error(w, verifyMethodTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	receivedPayload := new(models.Payload)

	defer r.Body.Close()
	decodePayloadErr := receivedPayload.FromJSON(r.Body)

	if decodePayloadErr != nil {
		http.Error(w, decodePayloadErr.Error(), http.StatusBadRequest)
		return
	}

	newRequestBody, newRequestErr := json.Marshal(receivedPayload)

	if newRequestErr != nil {
		http.Error(w, newRequestErr.Error(), http.StatusInternalServerError)
		return
	}

	request, err := http.NewRequest(http.MethodPut, "models-service:8081/pseudonym", bytes.NewBuffer(newRequestBody))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}

	modelResponse, modelErr := client.Do(request)

	if modelErr != nil {
		http.Error(w, modelErr.Error(), http.StatusFailedDependency)
		return
	}

	defer modelResponse.Body.Close()
	response := createResponse(
		"Added pseudonym successfully.",
		nil,
		modelResponse.Body)

	fmt.Printf("%v", response)
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	toJSONErr := response.ToJSON(w)

	if toJSONErr != nil {
		http.Error(w, toJSONErr.Error(), http.StatusInternalServerError)
		return
	}

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
