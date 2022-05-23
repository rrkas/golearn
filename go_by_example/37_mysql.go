package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	createTable()
	testConnection()
	insertData()
}

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", string("root:mysql@tcp(127.0.0.1:3306)/godb"))
}

func createTable() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	// prepare development
	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS sensordevice(
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			deviceid INTEGER,
			temperature REAL,
			humidity REAL,
			light_intensity REAL,
			pressure REAL
		)`,
	)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func testConnection() {
	// change database user and password
	db, err := getDB()
	if err != nil {
		panic(err)
	}
	err = db.Ping() // test connection
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")
	defer db.Close()
}

func insertData() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	// prepare development
	stmt, err := db.Prepare("INSERT INTO sensordevice(deviceid,temperature,humidity,light_intensity,pressure) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	fmt.Println("inserting data…")
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(3, 0.2*float64(i), 0.6*float64(i), 0.5*float64(i), 0.6*float64(i))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("done")
	defer db.Close()
}
