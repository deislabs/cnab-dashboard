package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Bundles int `json:"bundles"`
		}{
			Bundles: 25,
		}

		j, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err.Error())
		}

		fmt.Fprintln(w, string(j))
	})
	log.Fatal(http.ListenAndServe(":3031",nil))
}