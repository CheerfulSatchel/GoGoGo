package database

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

type customError struct {
	errorMessage string
}

type Pseudonym struct {
	tableName struct{} `sql:"pseudonym"`
	ID        int
	Username  string
}

type PseudonymDetails struct {
	tableName   struct{} `sql:"pseudonymdetails`
	HTMLURL     string
	ID          int
	Likes       int
	PseudonymID int
	Pseudonym   *Pseudonym
}

var pgdb *pg.DB

func init() {
	pgdb = connect()
	fmt.Println("Connected to database~")
	fmt.Println(pgdb)
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("PSEUDONYM_USER"),
		Password: os.Getenv("PSEUDONYM_PASSWORD"),
		Database: os.Getenv("PSEUDONYM_DATABASE"),
	})
}

func (e *customError) Error() string {
	return "ERROR: " + e.errorMessage
}

func CreateTables() error {
	for _, model := range []interface{}{&Pseudonym{}, &PseudonymDetails{}} {
		err := pgdb.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil {
			fmt.Println(err)
			return &customError{"Failure to create tables! "}
		}
	}

	fmt.Println("Successfully created tables!!")
	return nil
}

func InsertUserIntoTable(newEntry interface{}) error {
	if _, ok := newEntry.(*Pseudonym); ok {
		fmt.Printf("Received pseudonym entry type~~")
		err := pgdb.Insert(newEntry)
		return err
	} else {
		fmt.Printf("Received wrong entry type...")
		err := &customError{errorMessage: "Received wrong entry type..."}
		return err
	}
}

func InsertUserDetailIntoTable(newEntry interface{}) error {
	if _, ok := newEntry.(*PseudonymDetails); ok {
		fmt.Printf("Received pseudonym details entry type~~")
		err := pgdb.Insert(newEntry)
		return err
	} else {
		fmt.Printf("Received wrong entry type...")
		err := &customError{errorMessage: "Received wrong entry type..."}
		return err
	}
}

func Query(id int) (*PseudonymDetails, error) {
	returnPseudonymDetails := new(PseudonymDetails)

	err := pgdb.Model(returnPseudonymDetails).
		Relation("Pseudonym").
		Where("pseudonym_id = ?", id).
		Select()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return returnPseudonymDetails, nil

}
