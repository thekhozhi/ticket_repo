package postgres_connection

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	) 
func GoConnectingSql()(* sql.DB, error){
	db, err := sql.Open("postgres","host=localhost port=5432 user=the_khoji password=546944 database=ticket sslmode=disable")
	if err != nil{
		fmt.Println("Error while connecting to database")
	}
	return db, nil
}