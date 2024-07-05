package main

import (
	"encoding/json"
	"net/http"
)

type rental struct {
	store clerk
}

func (h rental) listHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(h.store.list())
	if err != nil {
		w.Write([]byte("Unable to encode shelf"))
	}
	// tj, err := json.Marshal(h.store.list())
	// if err != nil {
	// 	log.Fatalln("Unable to marshal shelf")
	// }

	// w.Write(tj)
}
