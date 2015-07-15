package main

import (
	"net/http"
	"fmt"
	"os"
)

var requests []string

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
	requests = append(requests, r.FormValue("text"))
	fmt.Fprintf(w, "Successfully requested %s!", r.FormValue("text"))
	fmt.Println("request for " + r.FormValue("text") + " saved successfully")
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("View request")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, `{"requests":[`)
	for i := 0; i < len(requests); i++ {
		fmt.Fprintf(w, `"%s"%s`, slice[i], i == len(requests) ? "" : ", ")
	}

	requests = nil
	
	fmt.Fprintf(w, `], "playlist":[`)
	fmt.Fprintf(w, `"XE-oMOEZ7Rc"`)
	fmt.Fprintf(w, `]}`)
}
