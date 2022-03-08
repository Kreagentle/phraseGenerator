package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type phrase struct{
	id int
	thought string
}

func phraseHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/phrases?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from motivatePhrases")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var thinks []phrase
	quantity := 0
	for rows.Next(){
		p := phrase{}
		err := rows.Scan(&p.id, &p.thought)
		if err != nil{
			fmt.Println(err)
			continue
		}
		thinks = append(thinks, p)
		quantity += 1
	}

	rand.Seed(time.Now().UnixNano())
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(thinks[rand.Intn(quantity)].thought); err != nil {
		return
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/phrases", phraseHandler)
	http.Handle("/",router)

	http.ListenAndServe(":8080", nil)
}
