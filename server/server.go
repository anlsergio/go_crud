package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// CreateUser adds a new user to the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// validates if the request body contains any information (JSON payload)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to get information from the request body"))
		return // empty return just to prevent the function execution from moving forward
	}

	var u user

	if err = json.Unmarshal(requestBody, &u); err != nil {
		w.Write([]byte("Failed converting the user information from JSON"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to open a database connection"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT - Prevents SQL injection exploitation in your query statement
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Failed to create SQL statement"))
	}
	defer statement.Close()

	insertion, err := statement.Exec(u.Name, u.Email)
	if err != nil {
		w.Write([]byte("Failed to execute the SQL statement"))
		return
	}

	userID, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to get the ID regarding the added user"))
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP STATUS CODE
	w.Write([]byte(fmt.Sprintf("The user has been successfuly added! User ID: %d", userID)))
}

// GetUsers fetch all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to open a database connection"))
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Failed to fetch users from the database"))
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			w.Write([]byte("Failed to parse the user information from the query set"))
			return
		}

		users = append(users, u)
	}

	w.WriteHeader(http.StatusOK)
	// convert the slice structure into JSON (this time specific for a HTTP response use case, 
	// JSON Marshal/Unmarshal doesn't work for HTTP use cases
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Failed converting users to JSON"))
		return
	}
}

// GetUser fetch a given user from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Capture Query params from the request URI
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed converting the query param from the request to integer"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to open a database connection"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Failed to fetch user from the database"))
		return
	}
	
	var u user
	if row.Next() {
		if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			w.Write([]byte("Failed to parse the user information from the query set"))
			return
		}
	}
	row.Close()

	if err := json.NewEncoder(w).Encode(u); err != nil {
		w.Write([]byte("Failed converting users to JSON"))
		return
	}
}

// UpdateUser update information from a given user in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed converting the query param from the request to integer"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to get information from the request body"))
		return
	}

	var u user
	if err := json.Unmarshal(requestBody, &u); err != nil {
		w.Write([]byte("Failed converting the user information from JSON"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to open a database connection"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Failed to create SQL statement"))
	}
	defer statement.Close()

	if _, err := statement.Exec(u.Name, u.Email, ID); err != nil {
		w.Write([]byte("Failed to execute the SQL statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteUser deletes a given user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed converting the query param from the request to integer"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to open a database connection"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Failed to create SQL statement"))
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Failed to execute the SQL statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}