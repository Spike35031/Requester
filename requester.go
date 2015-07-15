package main

import (
	"net/http"
	"fmt"
	"code.google.com/p/go-sqlite/go1/sqlite3"
)

func main() {
	conn, _ := sqlite3.Open(":memory:")
	defer conn.Close()
	conn.Exec("CREATE TABLE requests(id int, url VARCHAR(128)")
	http.HandleFunc("/request/", requestHandler)
	http.HandleFunc("/player/", playerHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if (r.FormValue("token") != "s2YxgysYXfGsbXLeH03La8pD") {
		fmt.Fprintf(w, "Bad token! Talk to Ben T.")
		return
	}
	conn, _ := sqlite3.Open(":memory:")
	defer conn.Close()
	conn.Exec("INSERT INTO requests(url) VALUES (" + r.FormValue("text"))
	fmt.Fprintf(w, "Successfully requested %s!", r.FormValue("text"))
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
