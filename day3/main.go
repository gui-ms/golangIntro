package main

import (
	//"fmt"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	//Adding driver for sqlite3
	//Here we use a blank identifier '_' for packages used indirectly
	_ "github.com/mattn/go-sqlite3"

	//Profile mode
	_ "net/http/pprof"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Calculate Fibonacci function
func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/cpu", CPUIntensiveEndpoint)
	//We put one of the entrypoints in another thread so that we can reach the other one
	go http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":6060", nil)
}

/*For the purposes of practice, we'll open a connection with a database in the controller,
though this is not recommended in a more professional environment.*/

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	//error handling:
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("Select * from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//Here we create a slice, which can be said to be kind of a dynamic array
	users := []User{}
	for rows.Next() {
		var u User
		//The '&' indicates direct access to the allocated memory address
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	//error handling and convert to json:
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	//error handling:
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := db.Exec(
		"INSERT INTO users (id, name , email) VALUES (?, ?, ?)",
		u.ID, u.Name, u.Email,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func CPUIntensiveEndpoint(w http.ResponseWriter, r *http.Request) {
	result := Fibonacci(60)
	//Convert to string and write
	w.Write([]byte(strconv.Itoa(result)))
}

func GenerateLargeString(size int) string {
	var buffer bytes.Buffer
	for i := range size {
		for j := range 100 {
			buffer.WriteString(strconv.Itoa(i + j*j))
		}
	}
	return buffer.String()
}

/* func GenerateLargeString(size int) string {
	var buffer bytes.Buffer
	buffer.Grow(size * 100)
	for i := 0; i < size; i++ {
		for j := 0; j < 100; j++ {
			buffer.WriteString(strconv.Itoa(i + j*j))
		}
	}
	return buffer.String()
}
*/
