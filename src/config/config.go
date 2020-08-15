package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

/* func GetDB() (db *sql.DB, err error) {

	dbDriver := "mssqldb"
	dbUser := "sa"
	dbPass := "pa$$w0rd"
	dbName := "BDCrudTest"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

} */

var server = "192.168.56.3"
var port = 1433
var user = "sa"
var password = "pa$$w0rd"
var database = "BDCrudTest"

//var db *sql.DB

func GetDB() (db *sql.DB, err error) {
	//var err error

	//Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;Trusted_Connection=True",
		server, user, password, port, database)

	fmt.Println(connString)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	// Close the database connection pool after program executes
	//defer db.Close()

	return

}
