package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database_service/database"
	"github.com/julienschmidt/httprouter"
)

type Response map[string]interface{}

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

func CreateTables(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	verifyMethodTypeErr := verifyMethodType(r.Method, http.MethodPut, &w)

	if verifyMethodTypeErr != nil {
		http.Error(w, verifyMethodTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	createTablesErr := database.CreateTables()

	if createTablesErr != nil {
		http.Error(w, createTablesErr.Error(), http.StatusInternalServerError)
		return
	}

	response := createResponse(
		fmt.Sprintf("Successfully created tables~"),
		nil,
		nil)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	toJSONErr := response.ToJSON(w)

	if toJSONErr != nil {
		http.Error(w, "Failure to encode JSON", http.StatusInternalServerError)
		return
	}
}
