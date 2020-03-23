package models

type Pseudonym struct {
	tableName struct{} `sql:"pseudonym"`
	ID        int      `json:"id"`
	Username  string   `json:"username"`
}

type PseudonymDetails struct {
	tableName   struct{}   `sql:"pseudonymdetails`
	HTMLURL     string     `json:"html_url"`
	ID          int        `json:"details_id"`
	Likes       int        `json:"likes"`
	PseudonymID int        `json:"-"`
	Pseudonym   *Pseudonym `json:"pseudonym"`
}
