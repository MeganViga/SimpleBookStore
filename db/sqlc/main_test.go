package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)


const (
	dbDriver = "postgres"
	dbSource ="postgres://root:secret@localhost:5432/book_store?sslmode=disable"
)
var testQueries *Queries
var testDB *sql.DB
func TestMain( m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver,dbSource)
	if err !=  nil{
		fmt.Println("Hello")
		log.Fatal(err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}