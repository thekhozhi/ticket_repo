package users

import (
	"database/sql"
	"errors"
	"fmt"
	pc "ticket/postgres_connection"
	"github.com/google/uuid"
)


var db, _ = pc.GoConnectingSql()
 



type User struct {
	id uuid.UUID
	first_name string
	last_name string
	email string
	phone string
	ticket_id uuid.UUID
}


var user = User{
id: uuid.New(),
first_name: "Khojiakbar",
last_name: "Rakhmatillayev",
email: "khojiakbar7@gmail.com",
phone: "+998 957777777",
ticket_id: uuid.New(),
}


func InsertUser() error{
	_, err := db.Exec(`INSERT INTO users values ($1, $2, $3, $4)`,  user.id, user.first_name, user.last_name, user.email, user.phone, user.ticket_id)
	if err != nil {
		fmt.Println("Error while inserting data into users!", err) 
	}
	return nil
}

func GetUserById(User) error{
	u := User{}
	id := ""
	row := db.QueryRow(`SELECT id, first_name, last_name, email, phone, ticket_id where id = $1`, id)
	err := row.Scan(&u.id, &u.first_name, &u.last_name, &u.email, &u.phone, &u.ticket_id)
	if err != nil{
		if !errors.Is(err, sql.ErrNoRows){
			fmt.Println("Error while scanning into users!", err)
	   }
    }
	fmt.Println(u)
	return nil
}

func ListOfUsers(User) error{
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil{
		fmt.Println("Error while selecting all users!", err)
	}
	allusers := []User{}

	for rows.Next(){
		u := User{}
		err := rows.Scan(&u.id, &u.first_name, &u.last_name, &u.email, &u.phone, &u.ticket_id)
		if err != nil{
			fmt.Println("Error while scanning into user struct!", err)
		}
		allusers = append(allusers, u)
	}
	return nil
}

func UpdateUserById (u User) error{
	id := ""
	_, err := db.Exec(`UPDATE users set id = $1, first_name = $2, last_name = $3, email = $4, phone = $5, ticket_id = $6  where id = $7`, u.id, u.first_name, u.last_name, u.email, u.phone, u.ticket_id,  id)
	if err != nil{
		fmt.Println("Error while updating users!", err)
	}
	return nil
}

func DeleteUser () error{
	id := ""
	_, err := db.Exec(`DELETE FROM users where id = $1`, id)
	if err != nil{
		fmt.Println("Error while deleting from users!", err)
	}
	return nil
}


