package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func verifyMethodType(givenMethod, acceptedMethod string, w *http.ResponseWriter) error {
	if strings.Compare(givenMethod, acceptedMethod) == 0 {
		return nil
	} else {
		responseMessage := fmt.Sprintf("Sorry, given method %s is not accepted.", givenMethod)
		return errors.New(responseMessage)
	}
}
