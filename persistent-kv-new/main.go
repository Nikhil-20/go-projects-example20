package main

import (
	"distributed-kv/store"
	"log"
	"net/http"
)

func main() {

	wal, err := store.NewWal("wal.log")
	if err != nil {
		log.Fatal(err)
	}

	st := store.NewStore(wal)

	http.HandleFunc("/set", st.HandleSet)
	http.HandleFunc("/get", st.HandleGet)

	log.Println("Server started at port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
