package tickets

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)


type Database struct {
	db *sql.DB
}

func New (db *sql.DB) Database{
	return Database{
		db:	db,
	}
}

type Ticket struct {
	id uuid.UUID
	from_city string
	to_city string
	date_of_flight string
}

var TickeT = Ticket{
 	id: uuid.New(),
	from_city: "Tashkent",
	to_city: "Dubai",
	date_of_flight: "2023-12-27 14:23:55",
}

func (d *Database)InsertTicket() error{
	_, err := d.db.Exec(`INSERT INTO tickets values ($1, $2, $3, $4)`, TickeT.id, TickeT.from_city, TickeT.to_city, TickeT.date_of_flight)
	if err != nil {
		fmt.Println("Error while inserting data into tickets!", err)
	}
	return nil
}

func (d *Database) GetTicketById(Ticket) error{
	t := Ticket{}
	id := "02a3d312-5433-444a-8018-2a62f1ecf065"
	row := d.db.QueryRow(`SELECT id, from_city, to_city, date_of_flight from tickets where id = $1`, id)
	err := row.Scan(&t.id, &t.from_city, &t.to_city, &t.date_of_flight)
	if err != nil{
		if !errors.Is(err, sql.ErrNoRows){
			fmt.Println("Error while scanning into ticket struct!", err)
		}
	}
	fmt.Println(t)
	return nil
}

func (d *Database) ListOfTickets(Ticket) error{
	rows, err := d.db.Query(`SELECT * FROM tickets`)
	if err != nil{
		fmt.Println("Error while selecting all tickets!", err)
	}
	alltickets := []Ticket{}

	for rows.Next(){
		t := Ticket{}
		err := rows.Scan(&t.id, &t.from_city, &t.to_city, &t.date_of_flight)
		if err != nil{
			fmt.Println("Error while scanning into ticket struct!", err)
		}
		alltickets = append(alltickets, t)
	}
	return nil
}

func (d *Database) UpdateTicketById (t Ticket) error{
	id := "8f0329fb-aa26-41bc-b396-91071e868062"
	_, err := d.db.Exec(`UPDATE tickets set id = $1, from_city = $2, to_city = $3, date_of_flight = $4 where id = &5`, t.id, t.from_city, t.to_city, t.date_of_flight, id)
	if err != nil{
		fmt.Println("Error while updating tickets!", err)
	}
	return nil
}

func (d *Database) DeleteTicket () error{
	id := "8f0329fb-aa26-41bc-b396-91071e868062"
	_, err := d.db.Exec(`DELETE FROM tickets where id = $1`, id)
	if err != nil{
		fmt.Println("Error while deleting from tickets!", err)
	}
	return nil
}