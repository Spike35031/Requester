package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

var requests []string

func main() {
	http.HandleFunc("/request/", requestHandler)
	http.HandleFunc("/player/", playerHandler)
	http.HandleFunc("/playlist/", playlistHandler)
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
		fmt.Fprintf(w, `"%s"`, requests[i])
		if (i != len(requests) - 1) {
			fmt.Fprintf(w, ", ")
		}
	}

	requests = nil

	playlist, _ := ioutil.ReadFile("playlist.json")
	fmt.Fprintf(w, `], "playlist":[%s]}`, playlist)	
}

func playlistHandler(w http.ResponseWriter, r *http.Request) {
	if (r.FormValue("token") != "fMirvlQ7N0JkKYb6wIH5VUTv") {
		fmt.Fprintf(w, "Bad token! Talk to Ben T.")
		return
	}
	f, _ := os.OpenFile("playlist.json", os.O_APPEND, 0666)
	defer f.Close()
	f.WriteString(`,\n"` + r.FormValue("text") + `"`)
	fmt.Println(r.FormValue("text") + " successfully added to playlist!")
}
