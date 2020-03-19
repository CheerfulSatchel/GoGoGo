package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ResponseError struct {
	ErrorCode int `json:"error_code"`
}

type NestedBoi struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type Response struct {
	Body         string     `json:"body"`
	ErrorMessage error      `json:"error_message"`
	Timestamp    string     `json:"timestamp"`
	Boi          *NestedBoi `json:"boi"`
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("Error code: %v", r.ErrorCode)
}

func GetResponse(isGoodResponse bool) ([]byte, error) {
	nestedBoi := &NestedBoi{
		ID:  24,
		URL: "github.com",
	}

	response := &Response{
		Body:      "Hello there bruh what is up",
		Timestamp: time.Now().Format(time.UnixDate),
		Boi:       nestedBoi,
	}

	if isGoodResponse {
		response.ErrorMessage = nil
	} else {
		response.ErrorMessage = &ResponseError{ErrorCode: 400}
	}

	bytes, err := json.MarshalIndent(response, "", " ")

	if err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}
