package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	DBConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DBConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT
	)`
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	getAllRow(DBConnection)
	//INSERT(DBConnection)
	//UPDATE(DBConnection)
	//DELETE(DBConnection)
}

func INSERT(DBConnection *sql.DB) {
	cmd := "INSERT INTO person (name, age) VALUES (?,?)"
	_, err := DBConnection.Exec(cmd, "Mike", 24)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ok")
}

func WHERE(DBConnection *sql.DB) {
	cmd := "SELECT * FROM person where age = ?"
	row := DBConnection.QueryRow(cmd, 1000)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)
}

func UPDATE(DBConnection *sql.DB) {
	cmd := "UPDATE person SET age = ? WHERE name = ?"
	_, err := DBConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}
}

func DELETE(DBConnection *sql.DB) {
	cmd := "DELETE FROM person WHERE name = ?"
	_, err := DBConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}

func SQLInjection(DBConnection *sql.DB) {
	tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X',100);"
	cmd := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DBConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err := rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

func getAllRow(DBConnection *sql.DB) {
	cmd := "SELECT * FROM person"
	rows, _ := DBConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
