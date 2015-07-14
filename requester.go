package main

import (
	"net/http"
	"fmt"
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"os"
)

func main() {
	os.Remove("database.db")
	conn, connErr := sqlite3.Open("database.db")
	if (connErr != nil) {
		fmt.Print(connErr)
		panic("Failed to open connection!")
	}
	defer conn.Close()
	tableErr := conn.Exec("CREATE TABLE IF NOT EXISTS requests(id INT AUTO_INCREMENT, url VARCHAR(128))")
	if (tableErr != nil) {
		fmt.Print(tableErr)
		panic("Failed to create table!")
	}
	http.HandleFunc("/request/", requestHandler)
	http.HandleFunc("/player/", playerHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if (r.FormValue("token") != "s2YxgysYXfGsbXLeH03La8pD") {
		fmt.Fprintf(w, "Bad token! Talk to Ben T.")
		return
	}
	conn, connErr := sqlite3.Open("database.db")
	if (connErr != nil) {
		fmt.Print(connErr)
		panic("Failed to open connection!")
	}
	defer conn.Close()
	err := conn.Exec("INSERT INTO requests VALUES ('', '" + r.FormValue("text") + "')")
	if (err != nil) {
		fmt.Print(err)
		panic("Failed to save")
	}
	fmt.Fprintf(w, "Successfully requested %s!", r.FormValue("text"))
	fmt.Println("request for " + r.FormValue("text") + " saved successfully")
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("View request")
	conn, err := sqlite3.Open("database.db")
	if (err != nil) {
		panic("Failed to open connection!")
	}
	defer conn.Close()
	sql := "SELECT url FROM requests"
	for s, err := conn.Query(sql); err == nil; err = s.Next() {
		var row string
		s.Scan(&row)
		fmt.Fprintf(w, "<br>%s", row)
		fmt.Println(row)
	}
}
