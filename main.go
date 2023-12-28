package main

import (
	"fmt"
	"ticket/postgres_connection"
	"ticket/tickets"
	_ "github.com/lib/pq"
 )
 

func main() {
	// CONNECTING GO POSTGRES
	db, err := postgres_connection.GoConnectingSql()
	if err != nil {
		fmt.Println("Error while connecting Go to SQL:", err)
		return
	}
	defer db.Close()

	dtb := tickets.New(db)

	//INSERTING ticket to sql

	err = dtb.InsertTicket()
	if err != nil {
		fmt.Println("Error inserting ticket:", err)
		return
	}
	fmt.Println("Ticket inserted successfully!")

	


}