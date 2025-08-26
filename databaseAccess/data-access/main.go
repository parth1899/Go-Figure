package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// creating a database handle
// note: The sql package must be used in conjunction with a database driver

var db *sql.DB // database handle: db variable of type *sql.DB, making it global for now

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	// capture connecion properties
	// Config

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// get database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// DSN string - data source name string
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	// albums, err := albumsByArtist("John Coltrane")
	albums, err := albumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	albID, err := addAlbum(Album{
		Title:  "Chahun Mai Ya Na",
		Artist: "Arijit Singh",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of the added album is %v\n", albID)

}

// QUERY FOR MULTIPLE ROWS
// function to query album of a given artist
// QUERY
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albums by artist %q : %v", name, err)
	}

	defer rows.Close() // Defer closing rows so that any resources it holds will be released when the function exits
	// defer is a keyword that schedules a function call to run after the surrounding function returns, no matter how it returns (normal return, error, or panic).
	// when you see defer rows.Close(), it means: "Don’t close the rows right now; close them automatically when the function exits."

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil

}

// QUERY FOR A SINGLE ROW
// QUERYROW
func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb, nil
}

// Use DB.Exec to execute an insert statement
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

// $env:DBPASS=""
// $env:DBUSER=""
