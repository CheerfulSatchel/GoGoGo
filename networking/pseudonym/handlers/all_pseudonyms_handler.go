package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AllPseudonymsHandler struct {
	logger *log.Logger
}

func NewAllPseudonymsHandler(logger *log.Logger) *AllPseudonymsHandler {
	return &AllPseudonymsHandler{logger}
}

func (allPseudonymsHandler *AllPseudonymsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("INCOMING REQUEST: %v", r.Method)
	io.WriteString(w, "Ur mum\n")
}
