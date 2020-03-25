package models

import (
	"encoding/json"
	"io"
)

type Payload struct {
	Username string `json:"user_name"`
	HTMLURL  string `json:"html_url"`
	Likes    int    `json:"likes" pg:",use_zero"`
}

func (payload *Payload) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	return decoder.Decode(payload)
}

func (payload *Payload) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", " ")

	return encoder.Encode(payload)
}
