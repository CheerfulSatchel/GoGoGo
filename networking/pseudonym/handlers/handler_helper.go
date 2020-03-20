package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func verifyMethodType(givenMethod, acceptedMethod string, response *Response, w *http.ResponseWriter) error {
	if strings.Compare(givenMethod, acceptedMethod) == 0 {
		return nil
	} else {
		response.Message = fmt.Sprintf("Sorry, given method %s is not accepted.", givenMethod)
		(*w).WriteHeader(http.StatusBadRequest)
		return errors.New(response.Message)
	}
}
