package postgres

/*
# setup

## go to psql console

## create a database

`CREATE DATABASE bookstore;`

## connect to the database

`\c bookstore`

## create table

CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

## Insert records
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

## view records

SELECT * FROM books;

## create a .env file

```
DB_USERNAME=dbusername
DB_PASSWORD=dbpassword
DB_HOST=localhost
DB_NAME=bookstore
DB_SSL_MODE=disable
DB_PORT=5432

#Alternative connection
#DB_CONN_STR=postgres://username:password@localhost/bookstore?sslmode=disable


```

# Credits
Todd McLeod => https://github.com/GoesToEleven/golang-web-dev

*/

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres driver import
)

var (
	db *sql.DB
)

func init() {
	var err error
	//load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Initialize db env variables
	dbUn := os.Getenv("DB_USERNAME")
	dbP := os.Getenv("DB_PASSWORD")
	dbH := os.Getenv("DB_HOST")
	dbN := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPortInt, _ := strconv.Atoi(dbPort)
	dbSSL := os.Getenv("DB_SSL_MODE")

	dbConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbH, dbPortInt, dbUn, dbP, dbN, dbSSL)

	// dbConn := os.Getenv("DB_CONN_STR") // Another way to put db connection
	db, err = sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Println("db connection failed")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("db connection failed")
		panic(err)
	}
	fmt.Println("You connected to your database")

}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":8080", nil)

	defer db.Close()
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, "did not select * from books", 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
