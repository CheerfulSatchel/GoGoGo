package models

import (
	"encoding/json"
	"io"
)

type Pseudonym struct {
	tableName struct{} `sql:"pseudonym"`
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	HTMLURL   string   `json:"html_url"`
	Likes     int      `json:"likes" pg:",use_zero"`
}

func (pseudonym *Pseudonym) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(pseudonym)
}
